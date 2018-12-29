import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

class Solution {
	public int reverse(int x) {
		boolean nativeInteger = (x < 0);
		char[] reservedInteger = ("" + x).toCharArray();
		char[] limit = (nativeInteger == false) ? ("" + 2147483647).toCharArray() :
			("" + -2147483648).toCharArray();
		
		int length = (nativeInteger == false) ? reservedInteger.length :
			reservedInteger.length - 1;
		int startIndex = (nativeInteger == false) ? 0 : 1;
		int j;
		int m;
		int n;
		char tmp;

		for (int i=startIndex; i<(length/2 + startIndex) ; i++){
			j = reservedInteger.length - i - 1 + startIndex;
			tmp = reservedInteger[i];
			reservedInteger[i] = reservedInteger[j];
			reservedInteger[j] = tmp;
		}
		
		if (reservedInteger.length > limit.length){
			return 0;
		} else if (reservedInteger.length == limit.length) {
			for (int i=startIndex; i<reservedInteger.length; i++){
				m = Character.getNumericValue(reservedInteger[i]);
				n = Character.getNumericValue(limit[i]);
				if (m > n) {
					return 0;
				}else if ( m < n ){
					break;
				}
			}
		}
		return Integer.parseInt(new String(reservedInteger));
	}
}

public class MainClass {
	public static void main(String[] args) throws IOException {
		BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
		String line;
		while ((line = in.readLine()) != null) {
			int x = Integer.parseInt(line);
			
			int ret = new Solution().reverse(x);
			
			String out = String.valueOf(ret);
			
			System.out.print(out);
		}
	}
}
