
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

// Java Code
class Main {
    public static void main(String[] args) {

long __start = System.nanoTime();


        Scanner sc = new Scanner(System.in);

        int T = sc.nextInt();
        while (T-- > 0) {
            long A = sc.nextLong();
            long B = sc.nextLong();
            System.out.println(A + B);
        }

        sc.close();
    

long __end = System.nanoTime();
System.out.println("\nTIME_MS=" + ((__end - __start) / 1_000_000));

}
}