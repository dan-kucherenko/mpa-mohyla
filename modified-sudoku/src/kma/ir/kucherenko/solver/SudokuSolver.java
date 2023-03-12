package kma.ir.kucherenko.solver;

public class SudokuSolver {
    /**
     * Solves a Sudoku puzzle using backtracking.
     *
     * @param puzzle the puzzle to solve
     * @param size   the size of the puzzle
     * @return true if a solution is found, false otherwise
     */
    public static boolean sudokuSolver(int[][] puzzle, int size) {
        int[] emptyCell = findEmptyCell(puzzle, size);
        if (emptyCell == null) {
            // If no empty cells are found, the puzzle is solved
            return true;
        }
        int row = emptyCell[0];
        int col = emptyCell[1];

        // If no empty cells are found, the puzzle is solved
        if (row == -1 && col == -1) {
            return true;
        }

        // Try different values for the empty cell
        for (int value = 1; value <= size; value++) {
            if (isValid(puzzle, row, col, value)) {
                puzzle[row][col] = value;
                if (sudokuSolver(puzzle, size)) {
                    return true;
                }
                puzzle[row][col] = 0;
            }
        }

        // If no value works for the empty cell, backtrack
        return false;
    }

    /**
     * Checks if a value is valid for a given cell in the puzzle.
     *
     * @param puzzle the puzzle to check
     * @param row    the row of the cell
     * @param col    the column of the cell
     * @param value  the value to check
     * @return true if the value is valid for the cell, false otherwise
     */
    private static boolean isValid(int[][] puzzle, int row, int col, int value) {
        int size = puzzle.length;
        int regionSize = (int) Math.sqrt(size);
        int regionRow = row / regionSize;
        int regionCol = col / regionSize;

        // Check row
        for (int j = 0; j < size; j++) {
            if (puzzle[row][j] == value) {
                return false;
            }
        }
        // Check column
        for (int i = 0; i < size; i++) {
            if (puzzle[i][col] == value)
                return false;
        }
        // Check region
        int startRow = regionRow * regionSize;
        int startCol = regionCol * regionSize;
        int endRow = Math.min(startRow + regionSize, size);
        int endCol = Math.min(startCol + regionSize, size);
        for (int i = startRow; i < endRow; i++) {
            for (int j = startCol; j < endCol; j++) {
                if (puzzle[i][j] == value)
                    return false;
            }
        }
        // Value is valid for the cell
        return true;
    }

    private static int[] findEmptyCell(int[][] puzzle, int size) {
        int[] firstCell = new int[]{-1, -1};
        // Find the cell in the top row with 1
        for (int j = 0; j < size; j++) {
            if (puzzle[0][j] == 1) {
                firstCell[0] = 0;
                firstCell[1] = j;
                break;
            }
        }
        // Start searching for empty cells from the first cell in the column with 1 in the top row
        for (int i = firstCell[0]; i < size; i++) {
            if (puzzle[i][firstCell[1]] == 0)
                return new int[]{i, firstCell[1]};
            // Move to the start and continue searching for empty cells
        }
        // if the column with 1 is full
        for (int startRow = 0; startRow < size; startRow++) {
            for (int startCol = 0; startCol < size; startCol++) {
                if (puzzle[startRow][startCol] == 0)
                    return new int[]{startRow, startCol};
            }
        }
        return null;
    }

    /**
     * Prints the Sudoku puzzle.
     *
     * @param puzzle the puzzle to print
     * @param size   the size of the puzzle
     */
    public static void printPuzzle(int[][] puzzle, int size) {
        for (int i = 0; i < size; i++) {
            for (int j = 0; j < size; j++) {
                System.out.print(puzzle[i][j] + " ");
            }
            System.out.println();
        }
    }
}
