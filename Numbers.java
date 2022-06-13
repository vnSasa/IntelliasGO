package numbers;

public class Numbers {
	public static void main(String[] args) {

//		determine which number we will check
		for (int i = 1; i < 150; i++) {
			int number = i;

			boolean flag = false;

			System.out.print("enter number is: " + number + " result: ");

			if (number == 1) {
				
//		the number 1 cannot be prime
				System.out.println("NO");
			
			} else if (number == 2) {
				
//		the number 2 is prime
				System.out.println("YES");
			
			} else if (number % 2 == 0) {
				
//		the rest of the even numbers cannot be prime
				System.out.println("NO");
			} else {

/* 		check the odd numbers to see if they are prime
		it is necessary to check whether the specified number 
		is divisible by any other in the range from 1 to itself 
		so that the fractional part was equal to 0
*/
				
				flag = true;
				int check = 3;

				for (int step = 0; step < number; step++) {

					if (check < number) {
						if (number % check == 0) {
							flag = false;
							break;
						} else {
							check += 2;
						}
					} else {
						break;
					}
				}

				if (!flag) {
					System.out.println("NO");
				} else {
					System.out.println("YES");
				}

			}
		}
	}

}