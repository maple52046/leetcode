import java.io.*;

class Solution {
    public boolean isPalindrome(int x) {
        int y, z;
        if (x < 0) {
            return false;
        }
        z = x;
        y = 0;
        while (x>0) {
            y = y*10 + x%10;
            x /= 10;
        }
        return (y == z);
    }
}

public class Main {
    public static String booleanToString(boolean input) {
        return input ? "True" : "False";
    }
    
    public static void main(String[] args) throws IOException {
        BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
        String line;
        while ((line = in.readLine()) != null) {
            int x = Integer.parseInt(line);
            
            boolean ret = new Solution().isPalindrome(x);
            
            String out = booleanToString(ret);
            
            System.out.print(out);
        }
    }
}