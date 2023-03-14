package kma.ir.kucherenko.solver;

public class SudokuSolver {
    /**
     * Solves a Sudoku using backtracking.
     *
     * @param sudoku the sudoku to solve
     * @param size   the size of the grid
     * @return true if a solution is found, false otherwise
     */
    public static boolean sudokuSolver(int[][] sudoku, int size) {
        int[] emptyCell = findEmptyCell(sudoku, size);
        if (emptyCell == null)
            // If no empty cells are found, the sudoku is solved
            return true;
        int row = emptyCell[0];
        int col = emptyCell[1];
        // Try different values for the empty cell
        for (int value = 1; value <= size; value++) {
            if (isValid(sudoku, row, col, value)) {
                sudoku[row][col] = value;
                if (sudokuSolver(sudoku, size))
                    return true;
                sudoku[row][col] = 0;
            }
        }
        // If no value works for the empty cell, backtrack
        return false;
    }

    /**
     * Checks if a value is valid for a given cell in the grid.
     *
     * @param grid the puzzle to check
     * @param row    the row of the cell
     * @param col    the column of the cell
     * @param value  the value to check
     * @return true if the value is valid for the cell, false otherwise
     */
    private static boolean isValid(int[][] grid, int row, int col, int value) {
        int size = grid.length;
        int regionSize = (int) Math.sqrt(size);
        int regionRow = row / regionSize;
        int regionCol = col / regionSize;

        // Check row
        for (int j = 0; j < size; j++) {
            if (grid[row][j] == value)
                return false;
        }
        // Check column
        for (int i = 0; i < size; i++) {
            if (grid[i][col] == value)
                return false;
        }
        // Check region
        int startRow = regionRow * regionSize;
        int startCol = regionCol * regionSize;
        int endRow = Math.min(startRow + regionSize, size);
        int endCol = Math.min(startCol + regionSize, size);
        for (int i = startRow; i < endRow; i++) {
            for (int j = startCol; j < endCol; j++) {
                if (grid[i][j] == value)
                    return false;
            }
        }
        // Value is valid for the cell
        return true;
    }

    private static int[] findEmptyCell(int[][] grid, int size) {
        int[] firstCell = new int[]{-1, -1};
        // Find the cell in the top row with 1
        for (int j = 0; j < size; j++) {
            if (grid[0][j] == 1) {
                firstCell[0] = 0;
                firstCell[1] = j;
                break;
            }
        }
        // Start searching for empty cells from the first cell in the column with 1 in the top row
        for (int i = firstCell[0]; i < size; i++) {
            if (grid[i][firstCell[1]] == 0)
                return new int[]{i, firstCell[1]};
            // Move to the start and continue searching for empty cells
        }
        // if the column with 1 is full
        for (int startRow = 0; startRow < size; startRow++) {
            for (int startCol = 0; startCol < size; startCol++) {
                if (grid[startRow][startCol] == 0)
                    return new int[]{startRow, startCol};
            }
        }
        return null;
    }

    /**
     * Prints the Sudoku puzzle.
     *
     * @param sudoku the sudoku to print
     * @param size   the size of the puzzle
     */
    public static void printPuzzle(int[][] sudoku, int size) {
        for (int i = 0; i < size; i++) {
            for (int j = 0; j < size; j++) {
                System.out.print(sudoku[i][j] + " ");
            }
            System.out.println();
        }
    }
}
