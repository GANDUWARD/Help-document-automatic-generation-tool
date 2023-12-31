import java.util.Scanner;

/**
 * @brief A simple class with a method to add two numbers.
 */
public class Example {

    /**
     * @brief Adds two numbers.
     * @param a The first number.
     * @param b The second number.
     * @return The sum of a and b.
     */
    public static int add(int a, int b) {
        return a + b;
    }

    /**
     * @brief The main method to demonstrate the functionality.
     * @param args Command-line arguments (not used in this example).
     */
    public static void main(String[] args) {
        // Create a Scanner object for user input
        Scanner scanner = new Scanner(System.in);

        // Prompt the user to enter two numbers
        System.out.print("Enter the first number: ");
        int num1 = scanner.nextInt();
        System.out.print("Enter the second number: ");
        int num2 = scanner.nextInt();

        // Calculate and display the sum
        int sum = add(num1, num2);
        System.out.println("The sum of " + num1 + " and " + num2 + " is: " + sum);

        // Close the Scanner
        scanner.close();
    }
}