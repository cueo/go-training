public class SwitchCase  {
	public static void main(String[] args) {
		
		int x = 10;
		int rem = x % 2;
		String result = switch(rem) {
			case 0 -> "Even";
			case 1 -> "Odd";
			default -> "WTH";
		};
		
		
		
		// String result = "";
// 		int rem = x % 2;
// 		switch(rem) {
// 			case 0:
// 				result = "Even";
// 				break;
// 			case 1:
// 				result = "Odd";
// 				break;
// 			default:
// 				result = "WTH";
// 		}
		System.out.println(result);
		
		
		// switch x :=10;  {
// 				case x % 2 == 0:
// 					println("x is even")
// 				case x % 2 	== 1:
// 					println("x is odd")
// 				default:
// 					println("huh")
// 			}
	}
}	