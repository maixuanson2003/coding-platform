"use client";
import React, { useState, useEffect, useRef } from "react";
import { useParams } from "next/navigation";
import {
  Code2,
  Clock,
  Zap,
  CheckCircle2,
  XCircle,
  Loader2,
  ChevronLeft,
  Play,
  Square,
} from "lucide-react";

export default function ProblemDetail() {
  const { id } = useParams();
  // Problem state
  const [problem, setProblem] = useState(null);
  const [loading, setLoading] = useState(true);

  // Code editor state
  const [code, setCode] = useState("");
  const [language, setLanguage] = useState("javascript");
  const [output, setOutput] = useState("");
  const [isRunning, setIsRunning] = useState(false);
  const [theme, setTheme] = useState("vs-dark");

  const editorRef = useRef(null);
  const monacoRef = useRef(null);
  const sseRef = useRef(null);

  const codeTemplates = {
    javascript: '// JavaScript Code\nconsole.log("Hello World!");',
    python: '# Python Code\nprint("Hello World!")',
    java: '// Java Code\nclass Main {\n    public static void main(String[] args) {\n        System.out.println("Hello World!");\n    }\n}',
    cpp: '// C++ Code\nint main() {\n    cout << "Hello World!" << endl;\n    return 0;\n}',
  };

  // Fetch problem detail
  useEffect(() => {
    const fetchProblem = async () => {
      setLoading(true);
      try {
        const response = await fetch(`http://localhost:8080/api/problem/${id}`);
        const result = await response.json();

        setProblem(result?.data);
        setCode(codeTemplates[language]);
      } catch (error) {
        console.error("Error fetching problem:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchProblem();
  }, [id]);

  // Initialize Monaco Editor
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
      if (sseRef.current) {
        sseRef.current.close();
        sseRef.current = null;
      }
    };
  }, []);

  // Update language
  useEffect(() => {
    if (monacoRef.current) {
      const model = monacoRef.current.getModel();
      window.monaco.editor.setModelLanguage(model, language);
    }
  }, [language]);

  // Update theme
  useEffect(() => {
    if (monacoRef.current) {
      window.monaco.editor.setTheme(theme);
    }
  }, [theme]);

  // SSE functions
  const stopSSE = (note) => {
    if (sseRef.current) {
      try {
        sseRef.current.close();
      } catch (e) {
        console.error(e);
      }
      sseRef.current = null;
    }
    setIsRunning(false);
    if (note) setOutput((prev) => (prev ? prev + "\n" : "") + note);
  };

  const startSSE = (submissionId) => {
    if (sseRef.current) {
      sseRef.current.close();
      sseRef.current = null;
    }

    const sseUrl = `http://localhost:8080/api/event/1/${id}/${submissionId}`;

    try {
      const es = new EventSource(sseUrl);
      sseRef.current = es;

      es.onmessage = (e) => {
        setOutput((prev) => (prev ? prev + "\n" : "") + e.data);
      };

      es.addEventListener("done", (e) => {
        try {
          const payload = e.data;
          setOutput((prev) => (prev ? prev + "\n" : "") + payload);
        } finally {
          stopSSE("✅ Hoàn thành");
        }
      });

      es.onerror = () => {
        setOutput(
          (prev) => (prev ? prev + "\n" : "") + "❌ SSE connection error"
        );
        stopSSE();
      };
    } catch (err) {
      setOutput("❌ Không thể mở kết nối SSE");
      setIsRunning(false);
    }
  };

  // Submit code
  const runCode = async () => {
    const userCode = monacoRef.current.getValue();

    setIsRunning(true);
    setOutput("🔄 Đang gửi code lên server...");

    try {
      const response = await fetch(
        `http://localhost:8080/api/submission?user_id=1&problem_id=${id}`,
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

      if (result && result?.data?.submission) {
        setOutput("🔁 Đang nhận luồng kết quả...");
        startSSE(result?.data?.submission);
      } else {
        setOutput("❌ Không nhận được submission ID");
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

  const getDifficultyColor = (difficulty) => {
    switch (difficulty?.toLowerCase()) {
      case "easy":
        return "text-green-400 bg-green-400/10 border-green-400/20";
      case "medium":
        return "text-yellow-400 bg-yellow-400/10 border-yellow-400/20";
      case "hard":
        return "text-red-400 bg-red-400/10 border-red-400/20";
      default:
        return "text-slate-400 bg-slate-400/10 border-slate-400/20";
    }
  };

  const languages = [
    { value: "javascript", label: "JavaScript" },
    { value: "python", label: "Python" },
    { value: "java", label: "Java" },
    { value: "cpp", label: "C++" },
  ];

  if (loading) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-slate-950 via-slate-900 to-slate-950 flex items-center justify-center">
        <Loader2 className="w-12 h-12 text-blue-400 animate-spin" />
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-950 via-slate-900 to-slate-950 text-white flex flex-col">
      {/* Header */}
      <header className="bg-slate-900/50 backdrop-blur-lg border-b border-slate-800 sticky top-0 z-50">
        <div className="px-6 py-4">
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-4">
              <button
                onClick={() => window.history.back()}
                className="flex items-center gap-2 text-slate-400 hover:text-white transition-colors"
              >
                <ChevronLeft className="w-5 h-5" />
                Back
              </button>
              <div className="w-px h-6 bg-slate-700"></div>
              <div className="flex items-center gap-2">
                <div className="w-8 h-8 bg-gradient-to-br from-blue-500 to-purple-600 rounded-lg flex items-center justify-center">
                  <Code2 className="w-5 h-5" />
                </div>
                <span className="text-xl font-bold">
                  Problem #{problem?.Id}
                </span>
              </div>
            </div>

            <div className="flex items-center gap-3">
              <button
                onClick={runCode}
                disabled={isRunning}
                className="flex items-center gap-2 px-6 py-2 bg-gradient-to-r from-green-500 to-emerald-600 rounded-lg font-semibold hover:shadow-lg hover:shadow-green-500/50 transition-all disabled:opacity-50 disabled:cursor-not-allowed"
              >
                {isRunning ? (
                  <>
                    <Loader2 className="w-4 h-4 animate-spin" />
                    Running...
                  </>
                ) : (
                  <>
                    <Play className="w-4 h-4" />
                    Submit
                  </>
                )}
              </button>
              {isRunning && (
                <button
                  onClick={() => stopSSE("⏹️ Đã huỷ bởi người dùng")}
                  className="flex items-center gap-2 px-4 py-2 bg-red-600 rounded-lg font-semibold hover:bg-red-700 transition-all"
                >
                  <Square className="w-4 h-4" />
                  Stop
                </button>
              )}
            </div>
          </div>
        </div>
      </header>

      {/* Main Content */}
      <div className="flex-1 flex overflow-hidden">
        {/* Left Panel - Problem Description */}
        <div className="w-1/2 border-r border-slate-800 overflow-y-auto">
          <div className="p-6">
            {/* Problem Title */}
            <div className="mb-6">
              <h1 className="text-3xl font-bold mb-4">{problem?.Title}</h1>
              <div className="flex flex-wrap items-center gap-3">
                <span
                  className={`px-3 py-1 rounded-full text-sm font-medium border ${getDifficultyColor(
                    problem?.Difficult
                  )}`}
                >
                  {problem?.Difficult}
                </span>
                <span className="px-3 py-1 rounded-full text-sm font-medium bg-slate-700/50 text-slate-300 border border-slate-600">
                  {problem?.Category}
                </span>
                <div className="flex items-center gap-1 text-slate-400 text-sm">
                  <Clock className="w-4 h-4" />
                  {problem?.TimeLimit}ms
                </div>
                <div className="flex items-center gap-1 text-slate-400 text-sm">
                  <Zap className="w-4 h-4" />
                  {problem?.MemoryLimit}MB
                </div>
              </div>
            </div>

            {/* Problem Content */}
            <div className="prose prose-invert max-w-none">
              <div className="text-slate-300 whitespace-pre-wrap leading-relaxed">
                {problem?.Content}
              </div>
            </div>

            {/* Stats */}
            <div className="mt-8 grid grid-cols-3 gap-4">
              <div className="bg-slate-800/50 rounded-lg p-4 border border-slate-700">
                <div className="text-slate-400 text-xs mb-1">Acceptance</div>
                <div className="text-2xl font-bold text-green-400">54.2%</div>
              </div>
              <div className="bg-slate-800/50 rounded-lg p-4 border border-slate-700">
                <div className="text-slate-400 text-xs mb-1">Submissions</div>
                <div className="text-2xl font-bold text-blue-400">123K</div>
              </div>
              <div className="bg-slate-800/50 rounded-lg p-4 border border-slate-700">
                <div className="text-slate-400 text-xs mb-1">Accepted</div>
                <div className="text-2xl font-bold text-purple-400">66.7K</div>
              </div>
            </div>
          </div>
        </div>

        {/* Right Panel - Code Editor */}
        <div className="w-1/2 flex flex-col">
          {/* Language Tabs */}
          <div className="bg-slate-900/50 border-b border-slate-800">
            <div className="flex gap-1 px-4">
              {languages.map((lang) => (
                <button
                  key={lang.value}
                  onClick={() => changeLanguage(lang.value)}
                  className={`px-4 py-3 font-medium transition-colors ${
                    language === lang.value
                      ? "bg-slate-800 text-white border-b-2 border-blue-500"
                      : "text-slate-400 hover:text-white hover:bg-slate-800/50"
                  }`}
                >
                  {lang.label}
                </button>
              ))}
            </div>
          </div>

          {/* Editor */}
          <div className="flex-1 overflow-hidden">
            <div ref={editorRef} className="w-full h-full" />
          </div>

          {/* Output Panel */}
          <div className="h-64 border-t border-slate-800 bg-slate-900/50">
            <div className="bg-slate-800/50 px-4 py-2 border-b border-slate-700">
              <h3 className="font-semibold text-sm flex items-center gap-2">
                <div
                  className={`w-2 h-2 rounded-full ${
                    isRunning ? "bg-yellow-400 animate-pulse" : "bg-slate-600"
                  }`}
                ></div>
                Output
              </h3>
            </div>
            <div className="p-4 overflow-auto h-[calc(100%-40px)] font-mono text-sm">
              {output ? (
                <pre className="whitespace-pre-wrap text-green-400">
                  {output}
                </pre>
              ) : (
                <div className="text-slate-500 text-center mt-16">
                  <div className="text-4xl mb-2">🚀</div>
                  <p>Run your code to see the output</p>
                </div>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
