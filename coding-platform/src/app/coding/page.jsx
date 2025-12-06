"use client";

import { useState, useEffect, useRef } from "react";

export default function CodeEditor() {
  const [code, setCode] = useState(
    '// Viết code của bạn ở đây\nconsole.log("Hello World!");'
  );
  const [language, setLanguage] = useState("javascript");
  const [output, setOutput] = useState("");
  const [theme, setTheme] = useState("vs-dark");
  const [isRunning, setIsRunning] = useState(false);
  const editorRef = useRef(null);
  const monacoRef = useRef(null);

  // ...existing code...
  const codeTemplates = {
    javascript: '// JavaScript Code\nconsole.log("Hello World!");',
    python: '# Python Code\nprint("Hello World!")',
    java: '// Java Code\nclass Main {\n    public static void main(String[] args) {\n        System.out.println("Hello World!");\n    }\n}',
    cpp: '// C++ Code\nint main() {\n    cout << "Hello World!" << endl;\n    return 0;\n}',
  };

  // New: SSE reference
  const sseRef = useRef(null);

  useEffect(() => {
    let editor;

    const initMonaco = async () => {
      if (!window.monaco) {
        const loaderScript = document.createElement("script");
        loaderScript.src =
          "https://cdnjs.cloudflare.com/ajax/libs/monaco-editor/0.45.0/min/vs/loader.min.js";
        document.body.appendChild(loaderScript);

        await new Promise((resolve) => {
          loaderScript.onload = resolve;
        });

        window.require.config({
          paths: {
            vs: "https://cdnjs.cloudflare.com/ajax/libs/monaco-editor/0.45.0/min/vs",
          },
        });

        await new Promise((resolve) => {
          window.require(["vs/editor/editor.main"], resolve);
        });
      }

      if (editorRef.current && !monacoRef.current) {
        editor = window.monaco.editor.create(editorRef.current, {
          value: code,
          language: language,
          theme: theme,
          automaticLayout: true,
          tabSize: 2,
          insertSpaces: true,
          minimap: { enabled: false },
          fontSize: 14,
          lineNumbers: "on",
          roundedSelection: false,
          scrollBeyondLastLine: false,
          readOnly: false,
        });

        monacoRef.current = editor;

        editor.onDidChangeModelContent(() => {
          setCode(editor.getValue());
        });
      }
    };

    initMonaco();

    return () => {
      if (monacoRef.current) {
        monacoRef.current.dispose();
        monacoRef.current = null;
      }
      // Close any active SSE when component unmounts
      if (sseRef.current) {
        sseRef.current.close();
        sseRef.current = null;
      }
    };
  }, []);

  useEffect(() => {
    if (monacoRef.current) {
      const model = monacoRef.current.getModel();
      window.monaco.editor.setModelLanguage(model, language);
    }
  }, [language]);

  useEffect(() => {
    if (monacoRef.current) {
      window.monaco.editor.setTheme(theme);
    }
  }, [theme]);

  // New: stop SSE
  const stopSSE = (note) => {
    if (sseRef.current) {
      try {
        sseRef.current.close();
      } catch (e) {
        /* ignore */
      }
      sseRef.current = null;
    }
    setIsRunning(false);
    if (note) setOutput((prev) => (prev ? prev + "\n" : "") + note);
  };

  // New: start SSE to stream updates for a submission
  const startSSE = (submissionId) => {
    // close previous if any
    if (sseRef.current) {
      sseRef.current.close();
      sseRef.current = null;
    }

    // Adjust endpoint to your backend's SSE stream path
    const sseUrl = `http://localhost:8080/api/event/1/1/${submissionId}`;
    console.log(sseUrl);

    try {
      const es = new EventSource(sseUrl);
      sseRef.current = es;

      // Basic messages
      es.onmessage = (e) => {
        // Append streaming chunk to output
        setOutput((prev) => (prev ? prev + "\n" : "") + e.data);
      };

      // Optional: custom event for final result (if server emits "done" or "result")
      es.addEventListener("done", (e) => {
        try {
          const payload = e.data;
          setOutput((prev) => (prev ? prev + "\n" : "") + payload);
        } finally {
          stopSSE("✅ Hoàn thành");
        }
      });

      es.onerror = (err) => {
        setOutput(
          (prev) => (prev ? prev + "\n" : "") + "❌ SSE connection error"
        );
        stopSSE();
      };
      console.log("run sse success");
    } catch (err) {
      setOutput("❌ Không thể mở kết nối SSE");
      setIsRunning(false);
    }
  };

  const runCode = async () => {
    const userCode = monacoRef.current.getValue();

    setIsRunning(true);
    setOutput("🔄 Đang gửi code lên server...");

    try {
      const response = await fetch(
        "http://localhost:8080/api/submission?user_id=1&problem_id=1",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            lang: language,
            code: userCode,
          }),
        }
      );

      const result = await response.json();
      console.log(result);
      console.log(result?.data?.submission);

      if (result && result?.data?.submission) {
        setOutput("🔁 Đang nhận luồng kết quả...");
        startSSE(result?.data?.submission);
      } else {
        setIsRunning(false);
      }
    } catch (err) {
      console.log(err);

      setOutput("❌ Lỗi kết nối API");
      setIsRunning(false);
    }
  };

  const changeLanguage = (newLang) => {
    setLanguage(newLang);
    const template = codeTemplates[newLang];
    if (template && monacoRef.current) {
      monacoRef.current.setValue(template);
      setCode(template);
    }
    setOutput("");
  };

  const clearCode = () => {
    if (monacoRef.current) {
      monacoRef.current.setValue("");
      setCode("");
      setOutput("");
    }
  };

  const downloadCode = () => {
    const extensions = {
      javascript: "js",
      python: "py",
      java: "java",
      cpp: "cpp",
    };

    const blob = new Blob([code], { type: "text/plain" });
    const url = URL.createObjectURL(blob);
    const a = document.createElement("a");
    a.href = url;
    a.download = `code.${extensions[language] || "txt"}`;
    a.click();
    URL.revokeObjectURL(url);
  };

  // ...existing code...
  const languages = [
    { value: "javascript", label: "🟨 JavaScript", icon: "JS" },
    { value: "python", label: "🐍 Python", icon: "PY" },
    { value: "java", label: "☕ Java", icon: "JAVA" },
    { value: "cpp", label: "⚡ C++", icon: "C++" },
  ];

  const themes = [
    { value: "vs-dark", label: "🌙 Dark" },
    { value: "vs", label: "☀️ Light" },
    { value: "hc-black", label: "🎯 High Contrast" },
  ];

  return (
    <div className="min-h-screen bg-gray-900 text-white">
      {/* Header */}
      <header className="bg-gray-800 border-b border-gray-700 px-6 py-4">
        <div className="max-w-7xl mx-auto flex items-center justify-between flex-wrap gap-4">
          <div className="flex items-center gap-3">
            <div className="w-10 h-10 bg-gradient-to-br from-blue-500 to-purple-600 rounded-lg flex items-center justify-center">
              <span className="text-xl font-bold">{"</>"}</span>
            </div>
            <h1 className="text-2xl font-bold">Multi-Language Code Editor</h1>
          </div>

          <div className="flex items-center gap-3">
            <select
              value={theme}
              onChange={(e) => setTheme(e.target.value)}
              className="bg-gray-700 text-white px-3 py-2 rounded-lg border border-gray-600 focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm"
            >
              {themes.map((t) => (
                <option key={t.value} value={t.value}>
                  {t.label}
                </option>
              ))}
            </select>
          </div>
        </div>
      </header>

      {/* Language Tabs */}
      <div className="bg-gray-800 border-b border-gray-700">
        <div className="max-w-7xl mx-auto px-6">
          <div className="flex gap-2 overflow-x-auto">
            {languages.map((lang) => (
              <button
                key={lang.value}
                onClick={() => changeLanguage(lang.value)}
                className={`px-6 py-3 font-semibold transition-colors whitespace-nowrap ${
                  language === lang.value
                    ? "bg-gray-700 text-white border-b-2 border-blue-500"
                    : "text-gray-400 hover:text-white hover:bg-gray-750"
                }`}
              >
                {lang.label}
              </button>
            ))}
          </div>
        </div>
      </div>

      {/* Main Content */}
      <main className="max-w-7xl mx-auto p-6">
        <div className="grid lg:grid-cols-2 gap-6">
          {/* Editor Panel */}
          <div className="bg-gray-800 rounded-lg overflow-hidden shadow-xl">
            <div className="bg-gray-700 px-4 py-3 flex items-center justify-between border-b border-gray-600">
              <h2 className="font-semibold flex items-center gap-2">
                <span className="w-3 h-3 bg-green-500 rounded-full"></span>
                Trình soạn thảo
              </h2>
              <div className="flex gap-2">
                <button
                  onClick={clearCode}
                  className="px-3 py-1 bg-gray-600 hover:bg-gray-500 rounded text-sm transition-colors"
                  title="Xóa code"
                >
                  🗑️ Xóa
                </button>
                <button
                  onClick={downloadCode}
                  className="px-3 py-1 bg-gray-600 hover:bg-gray-500 rounded text-sm transition-colors"
                  title="Tải xuống"
                >
                  💾 Tải
                </button>
              </div>
            </div>
            <div
              ref={editorRef}
              className="w-full"
              style={{ height: "550px" }}
            />
          </div>

          {/* Output Panel */}
          <div className="bg-gray-800 rounded-lg overflow-hidden flex flex-col shadow-xl">
            <div className="bg-gray-700 px-4 py-3 flex items-center justify-between border-b border-gray-600">
              <h2 className="font-semibold flex items-center gap-2">
                <span className="w-3 h-3 bg-blue-500 rounded-full"></span>
                Kết quả
              </h2>
              <div className="flex items-center gap-2">
                <button
                  onClick={runCode}
                  disabled={isRunning}
                  className="px-5 py-2 bg-green-600 hover:bg-green-700 disabled:bg-gray-600 disabled:cursor-not-allowed rounded font-semibold transition-colors flex items-center gap-2"
                >
                  {isRunning ? "⏳ Đang chạy..." : "▶️ Chạy code"}
                </button>
                {isRunning && (
                  <button
                    onClick={() => stopSSE("⏹️ Đã huỷ bởi người dùng")}
                    className="px-3 py-2 bg-red-600 hover:bg-red-700 rounded font-semibold transition-colors text-sm"
                  >
                    ⏹️ Dừng
                  </button>
                )}
              </div>
            </div>
            <div
              className="flex-1 p-4 font-mono text-sm overflow-auto bg-gray-900"
              style={{ height: "550px" }}
            >
              {output ? (
                <pre className="whitespace-pre-wrap text-green-400">
                  {output}
                </pre>
              ) : (
                <div className="text-gray-500 text-center mt-32">
                  <div className="text-6xl mb-4">🚀</div>
                  <p className="text-lg">Nhấn "Chạy code" để xem kết quả</p>
                  <p className="text-xs mt-2">
                    Hỗ trợ: JavaScript, Python, Java, C++
                  </p>
                </div>
              )}
            </div>
          </div>
        </div>

        {/* Info Cards */}
        <div className="grid md:grid-cols-4 gap-4 mt-6">
          <div className="bg-gray-800 p-5 rounded-lg border border-gray-700 hover:border-blue-500 transition-colors">
            <div className="text-3xl mb-2">🟨</div>
            <h3 className="font-semibold mb-2 text-yellow-400">JavaScript</h3>
            <p className="text-sm text-gray-300">
              Chạy trực tiếp trong trình duyệt với eval()
            </p>
          </div>

          <div className="bg-gray-800 p-5 rounded-lg border border-gray-700 hover:border-blue-500 transition-colors">
            <div className="text-3xl mb-2">🐍</div>
            <h3 className="font-semibold mb-2 text-blue-400">Python</h3>
            <p className="text-sm text-gray-300">
              Chạy qua Piston API với Python 3
            </p>
          </div>

          <div className="bg-gray-800 p-5 rounded-lg border border-gray-700 hover:border-blue-500 transition-colors">
            <div className="text-3xl mb-2">☕</div>
            <h3 className="font-semibold mb-2 text-orange-400">Java</h3>
            <p className="text-sm text-gray-300">
              Compile và chạy code Java online
            </p>
          </div>

          <div className="bg-gray-800 p-5 rounded-lg border border-gray-700 hover:border-blue-500 transition-colors">
            <div className="text-3xl mb-2">⚡</div>
            <h3 className="font-semibold mb-2 text-purple-400">C++</h3>
            <p className="text-sm text-gray-300">
              Compile và chạy với GCC compiler
            </p>
          </div>
        </div>

        {/* Features */}
        <div className="mt-6 bg-gradient-to-r from-blue-900/30 to-purple-900/30 rounded-lg p-6 border border-blue-500/30">
          <h3 className="text-xl font-bold mb-4 flex items-center gap-2">
            ✨ Tính năng nổi bật
          </h3>
          <div className="grid md:grid-cols-3 gap-4 text-sm">
            <div className="flex items-start gap-2">
              <span className="text-green-400">✓</span>
              <div>
                <div className="font-semibold">Syntax Highlighting</div>
                <div className="text-gray-400">
                  Highlight code theo từng ngôn ngữ
                </div>
              </div>
            </div>
            <div className="flex items-start gap-2">
              <span className="text-green-400">✓</span>
              <div>
                <div className="font-semibold">Auto-completion</div>
                <div className="text-gray-400">Gợi ý code thông minh</div>
              </div>
            </div>
            <div className="flex items-start gap-2">
              <span className="text-green-400">✓</span>
              <div>
                <div className="font-semibold">Online Execution</div>
                <div className="text-gray-400">
                  Chạy code ngay trên trình duyệt
                </div>
              </div>
            </div>
          </div>
        </div>
      </main>

      {/* Footer */}
      <footer className="mt-8 py-6 border-t border-gray-800 text-center text-gray-500 text-sm">
        <p>
          Powered by Monaco Editor & Piston API • Built with Next.js & React
        </p>
      </footer>
    </div>
  );
}
