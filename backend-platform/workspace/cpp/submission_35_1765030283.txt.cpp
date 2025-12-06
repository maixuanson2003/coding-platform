
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

int main() {
    long __mem_before = getMemoryKB();
    auto __start = chrono::high_resolution_clock::now();

    int n;
    cin >> n;

    vector<long long> nums(n);
    for (int i = 0; i < n; i++) cin >> nums[i];

    long long target;
    cin >> target;

    unordered_map<long long, int> mp;  // value -> index

    for (int i = 0; i < n; i++) {
        long long need = target - nums[i];

        if (mp.count(need)) {
            cout << mp[need] << " " << i;
            
        }

        mp[nums[i]] = i;
    }

    

    auto __end = chrono::high_resolution_clock::now();
    long __mem_after = getMemoryKB();
    cout << "\nTIME_MS=" << chrono::duration_cast<chrono::milliseconds>(__end - __start).count();
    cout << "\nMEMORY_KB=" << (__mem_after - __mem_before);

    return 0;
}