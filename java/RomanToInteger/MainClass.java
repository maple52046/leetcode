import java.io.*;
import javax.json.JsonArray;

class Solution {
    public int romanToInt(String s) {
        HashMap<Character, Integer> romanMap = new HashMap<Character, Integer>();
        int output = 0;
        int previous = 0;
        int current = 0;
        
        romanMap.put('I', 1);
        romanMap.put('V', 5);
        romanMap.put('X', 10);
        romanMap.put('L', 50);
        romanMap.put('C', 100);
        romanMap.put('D', 500);
        romanMap.put('M', 1000);
        
        for(int i=(s.length()-1); i>=0; i--){
            current = romanMap.get(s.charAt(i));
            if (previous != 0 && previous > current) {
                current *= -1;
                previous = 0;
            } else {
                previous = current;
            }
            output += current;
        }
        return output;
    }
}

public class MainClass {
    public static String stringToString(String input) {
        return JsonArray.readFrom("[" + input + "]").get(0).asString();
    }
    
    public static void main(String[] args) throws IOException {
        BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
        String line;
        while ((line = in.readLine()) != null) {
            String s = stringToString(line);
            
            int ret = new Solution().romanToInt(s);
            
            String out = String.valueOf(ret);
            
            System.out.print(out);
        }
    }
}