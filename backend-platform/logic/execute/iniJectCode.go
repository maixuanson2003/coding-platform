package execute

import (
	"strings"
)

func injectIntoMain(code, start, end string) string {
	// Tìm "int main"
	mainIdx := strings.Index(code, "int main")
	if mainIdx < 0 {
		return code
	}

	braceStart := strings.Index(code[mainIdx:], "{")
	if braceStart < 0 {
		return code
	}
	braceStart += mainIdx + 1

	braceEnd := strings.LastIndex(code, "}")
	if braceEnd < 0 {
		return code
	}

	body := code[braceStart:braceEnd]

	body = strings.ReplaceAll(body, "return 0;", "")
	body = strings.ReplaceAll(body, "return;", "")

	newBody := start + body + end + "\n    return 0;\n"

	return code[:braceStart] + newBody + code[braceEnd:]
}

func injectIntoJavaMain(code, start, end string) string {
	mainIdx := strings.Index(code, "public static void main")
	if mainIdx < 0 {
		return code
	}

	// tìm dấu '{' của hàm main
	openIdx := strings.Index(code[mainIdx:], "{")
	if openIdx < 0 {
		return code
	}
	openIdx += mainIdx

	// đếm block {} để tìm đúng block của main()
	level := 1
	bodyEnd := -1

	for i := openIdx + 1; i < len(code); i++ {
		switch code[i] {
		case '{':
			level++
		case '}':
			level--
			if level == 0 { // ← đúng dấu '}' thuộc về main
				bodyEnd = i
				goto Inject
			}
		}
	}

Inject:
	if bodyEnd < 0 {
		return code
	}

	// Inject vào đúng bên trong main()
	return code[:openIdx+1] +
		"\n" + start + "\n" +
		code[openIdx+1:bodyEnd] +
		"\n" + end + "\n" +
		code[bodyEnd:]
}

func injectIntoMainJS(code, start, end string) string {
	return start + code + end
}

var InjectStruct = map[string]func(code string) string{

	// ======================== C++ ========================
	"cpp": func(code string) string {

		header := `
#include <bits/stdc++.h>
#include <sys/resource.h>
#include <chrono>
using namespace std;

// FAST IO
static const auto ___fast_io = [](){
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    return 0;
}();

long getMemoryKB() {
    struct rusage r;
    getrusage(RUSAGE_SELF, &r);
    return r.ru_maxrss;
}
`

		metricStart := `
    long __mem_before = getMemoryKB();
    auto __start = chrono::high_resolution_clock::now();
`
		metricEnd := `
    auto __end = chrono::high_resolution_clock::now();
    long __mem_after = getMemoryKB();
    cout << "\nTIME_MS=" << chrono::duration_cast<chrono::milliseconds>(__end - __start).count();
    cout << "\nMEMORY_KB=" << (__mem_after - __mem_before);
`

		// nếu user chưa có include fastio thì bổ sung
		if !strings.Contains(code, "FAST IO") {
			code = header + "\n" + code
		}

		// inject metric vào trong main user
		return injectIntoMain(code, metricStart, metricEnd)
	},

	// ======================== JAVA ========================
	"java": func(code string) string {

		header := `
import java.io.*;
import java.util.*;

class FastScanner {
    private final InputStream in = System.in;
    private final byte[] buffer = new byte[1<<16];
    private int ptr = 0, len = 0;

    private int read() throws IOException {
        if (ptr >= len) {
            len = in.read(buffer);
            ptr = 0;
            if (len <= 0) return -1;
        }
        return buffer[ptr++];
    }

    int nextInt() throws IOException {
        int c;
        while((c = read()) <= ' ') if(c == -1) return -1;
        int sign = 1;
        if (c == '-') { sign = -1; c = read(); }
        int val = c - '0';
        while((c = read()) > ' ') val = val * 10 + (c - '0');
        return val * sign;
    }

    String next() throws IOException {
        int c;
        while((c = read()) <= ' ') if(c == -1) return null;
        StringBuilder sb = new StringBuilder();
        sb.append((char)c);
        while((c = read()) > ' ') sb.append((char)c);
        return sb.toString();
    }
}
`

		metricStart := `
long __start = System.nanoTime();
`
		metricEnd := `
long __end = System.nanoTime();
System.out.println("\nTIME_MS=" + ((__end - __start) / 1_000_000));
`

		code = header + "\n" + code

		// inject vào trong "public static void main"
		return injectIntoJavaMain(code, metricStart, metricEnd)
	},

	// ======================== PYTHON ========================
	"python": func(code string) string {

		header := `
import sys
import time
import resource

input = sys.stdin.readline

def mem_kb():
    return resource.getrusage(resource.RUSAGE_SELF).ru_maxrss
`

		metricStart := `
__mem_before = mem_kb()
__start = time.perf_counter()
`

		metricEnd := `
__end = time.perf_counter()
__mem_after = mem_kb()

print(f"\nTIME_MS={(__end-__start)*1000:.3f}")
print(f"MEMORY_KB={__mem_after - __mem_before}")
`

		// Ghép theo thứ tự đúng
		final := header + "\n" + metricStart + "\n" + code + "\n" + metricEnd

		return final
	},

	// ======================== JAVASCRIPT ========================
	"js": func(code string) string {

		header := `
const fs = require("fs");
const { performance } = require("perf_hooks");

const input = fs.readFileSync(0, "utf8").trim().split(/\s+/);
let idx = 0;
const next = () => input[idx++];

function memoryKB() {
    return process.memoryUsage().rss / 1024;
}
`

		metricStart := `
const __mem_before = memoryKB();
const __start = performance.now();
`
		metricEnd := `
const __end = performance.now();
const __mem_after = memoryKB();

console.log("\nTIME_MS=" + (__end - __start));
console.log("MEMORY_KB=" + (__mem_after - __mem_before));
`

		if !strings.Contains(code, "memoryKB") {
			code = header + "\n" + code
		}

		return injectIntoMainJS(code, metricStart, metricEnd)
	},
}
