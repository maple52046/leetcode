class Solution {
    public boolean isValid(String s) {
        char[] stack = new char[s.length()];
        int top = -1;
        HashMap<Character, Character> symbolMap = new HashMap<Character, Character>();
        
        symbolMap.put('(', ')');
        symbolMap.put('{', '}');
        symbolMap.put('[', ']');
        
        if (s.length()%2 > 0) return false;
        
        for (char symbol : s.toCharArray()) {
            System.out.println("symbol: " + symbol);
            switch (symbol) {
                case '(':
                case '{':
                case '[':
                    top++;
                    stack[top] = symbolMap.get(symbol);
                    break;
                default:
                    if (top <0) return false;
                    if (symbol != stack[top]) return false;
                    top--;
            }
        }
        
        return !(top > 0);
    }
}

public class MainClass {
    public static String stringToString(String input) {
        return JsonArray.readFrom("[" + input + "]").get(0).asString();
    }
    
    public static String booleanToString(boolean input) {
        return input ? "True" : "False";
    }
    
    public static void main(String[] args) throws IOException {
        BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
        String line;
        while ((line = in.readLine()) != null) {
            String s = stringToString(line);
            
            boolean ret = new Solution().isValid(s);
            
            String out = booleanToString(ret);
            
            System.out.print(out);
        }
    }
}