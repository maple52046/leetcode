class Solution {
    public String longestCommonPrefix(String[] strs) {
        char p;
        int prefixLength = 0;
        boolean stop = false;
        
        if (strs.length == 0) return "";
        if (strs.length == 1) return strs[0];
        
        for (int j=0; stop==false && j<strs[0].length(); j++) {
            p = strs[0].charAt(j);
            for (int i=1; i<strs.length; i++) {
                if (j >= strs[i].length()) {
                    stop = true;
                    break;
                }
                
                if ( p != strs[i].charAt(j) ){
                    stop = true;
                    break;
                }
            }
            
            if (stop == false) prefixLength += 1;
        }
        
        return strs[0].substring(0, prefixLength);
        
    }
}

public class MainClass {
    public static String[] stringToStringArray(String line) {
        JsonArray jsonArray = JsonArray.readFrom(line);
        String[] arr = new String[jsonArray.size()];
        for (int i = 0; i < arr.length; i++) {
          arr[i] = jsonArray.get(i).asString();
        }
        return arr;
    }
    
    public static void main(String[] args) throws IOException {
        BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
        String line;
        while ((line = in.readLine()) != null) {
            String[] strs = stringToStringArray(line);
            
            String ret = new Solution().longestCommonPrefix(strs);
            
            String out = (ret);
            
            System.out.print(out);
        }
    }
}