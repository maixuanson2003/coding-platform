
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

// C++ Code
int main() {
    long __mem_before = getMemoryKB();
    auto __start = chrono::high_resolution_clock::now();

    string s;
    cin >> s;
    
    string clean = "";
    for (char c : s) {
        if (isalnum(c)) clean += tolower(c);
    }

    // Two pointers kiểm tra palindrome
    int l = 0, r = clean.size() - 1;
    bool isPalin = true;

    while (l < r) {
        if (clean[l] != clean[r]) {
            isPalin = false;
            break;
        }
        l++;
        r--;
    }

    cout << (isPalin ? "true" : "false");
    

    auto __end = chrono::high_resolution_clock::now();
    long __mem_after = getMemoryKB();
    cout << "\nTIME_MS=" << chrono::duration_cast<chrono::milliseconds>(__end - __start).count();
    cout << "\nMEMORY_KB=" << (__mem_after - __mem_before);

    return 0;
}