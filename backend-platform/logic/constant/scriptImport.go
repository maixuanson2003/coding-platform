package constant

var (
	CppTemplate = `g++ -O2 -std=c++17 {{nameFile}} -o solution
./solution << EOF
{{input}}
EOF`
	CsharpTemplate = `dotnet run --project /app << EOF
{{input}}
EOF`
	JavaTemplate = `javac {{nameFile}}
java Main << EOF
{{input}}
EOF`
	JsTemplate = `node {{nameFile}} << EOF
{{input}}
EOF`
	PythonTemplate = `python3 {{nameFile}} << EOF
{{input}}
EOF`
)
var DomainExtension = map[string]string{
	"cpp":    "cpp",
	"Csharp": "cs",
	"java":   "java",
	"js":     "js",
	"python": "py",
}
var WorkerFolderPath = map[string]string{
	"cpp":    "workspace\\cpp",
	"Csharp": "workspace\\csharp",
	"java":   "workspace\\java",
	"js":     "workspace\\js",
	"python": "workspace\\python",
}
var FastIOHeader = map[string]string{
	"cpp": `#include <bits/stdc++.h>
using namespace std;
#define fastio ios::sync_with_stdio(false); cin.tie(nullptr);

`,
	"python": `import sys, time, resource

`,
	"java": `import java.io.*; 
import java.util.*;

`,
	"js": `"use strict";

`,
	"csharp": `using System;
using System.Diagnostics;

`,
}

var MetricHeader = map[string]string{
	"cpp": `
int main() {
    fastio
    auto __start = chrono::high_resolution_clock::now();
    long __mem_before = 0;
`,
	"python": `
__start = time.perf_counter()
__mem_before = resource.getrusage(resource.RUSAGE_SELF).ru_maxrss
`,
	"java": `
public class Main {
    public static void main(String[] args) throws Exception {
        long __start = System.nanoTime();
        long __mem_before = Runtime.getRuntime().totalMemory() - Runtime.getRuntime().freeMemory();
`,
	"js": `
const __start = performance.now();
const __mem_before = process.memoryUsage().rss;
`,
	"csharp": `
class Program {
    static void Main(string[] args) {
        var __watch = Stopwatch.StartNew();
        long __mem_before = Process.GetCurrentProcess().WorkingSet64;
`,
}

var MetricFooter = map[string]string{
	"cpp": `
    auto __end = chrono::high_resolution_clock::now();
    cout << "\nTIME_MS=" 
         << chrono::duration_cast<chrono::milliseconds>(__end - __start).count();
    return 0;
}
`,
	"python": `
__end = time.perf_counter()
__mem_after = resource.getrusage(resource.RUSAGE_SELF).ru_maxrss
print(f"\nTIME_MS={(__end-__start)*1000:.3f}")
`,
	"java": `
        long __end = System.nanoTime();
        long __mem_after = Runtime.getRuntime().totalMemory() - Runtime.getRuntime().freeMemory();
        System.out.println("\nTIME_MS=" + ( __end - __start ) / 1_000_000);
    }
}
`,
	"js": `
const __end = performance.now();
console.log("\\nTIME_MS=" + (__end - __start).toFixed(3));
`,
	"csharp": `
        __watch.Stop();
        long __mem_after = Process.GetCurrentProcess().WorkingSet64;
        Console.WriteLine($"\\nTIME_MS={__watch.ElapsedMilliseconds}");
    }
}
`,
}
