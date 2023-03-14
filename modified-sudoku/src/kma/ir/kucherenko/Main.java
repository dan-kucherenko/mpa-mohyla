package kma.ir.kucherenko;

import kma.ir.kucherenko.solver.SudokuSolver;

import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        // Get size and first row of the puzzle from the user
        System.out.print("Enter size of puzzle: ");
        int size = scanner.nextInt();
        int[][] sudokuGrid = new int[size][size];
        System.out.print("Enter first row of puzzle: ");
        for (int j = 0; j < size; j++) {
            sudokuGrid[0][j] = scanner.nextInt();
        }
        boolean solutionFound = SudokuSolver.sudokuSolver(sudokuGrid, size);
        if (solutionFound) {
            System.out.println("Solution:");
            SudokuSolver.printPuzzle(sudokuGrid, size);
        } else {
            System.out.println("No solution found.");
        }
    }
}