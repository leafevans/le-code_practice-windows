using System;
using System.Collections.Generic;
using System.Linq;
using LPQ.Domain;

namespace LPQ.Algorithm
{
    public class SimplexSolver
    {
        private double[,] _tableau = null!;
        private int _numVars;
        private int _numConstraints;

        public Solution Compute(LPQModel model)
        {
            try
            {
                Standardize(model);
                while (true)
                {
                    int enteringCol = FindEnteringColumn();
                    if (enteringCol == -1)
                    {
                        var sol = ExtractSolution(model);
                        if (
                            string.Equals(
                                model.ObjectiveType,
                                "Min",
                                StringComparison.OrdinalIgnoreCase
                            )
                        )
                        {
                            sol.MaxValue = -sol.MaxValue;
                        }
                        return sol;
                    }
                    int leavingRow = FindLeavingRow(enteringCol);
                    if (leavingRow == -1)
                    {
                        return new Solution { IsOptimal = false, Message = "问题无界 (Unbounded)" };
                    }
                    Pivot(leavingRow, enteringCol);
                }
            }
            catch (Exception ex)
            {
                return new Solution { IsOptimal = false, Message = "计算出错: " + ex.Message };
            }
        }

        private void Standardize(LPQModel model)
        {
            _numVars = model.Variables.Count;
            var processedConstraints = new List<(Dictionary<string, double> Coeffs, double Rhs)>();
            foreach (var cons in model.Constraints)
            {
                if (cons.Operator == "<=")
                {
                    processedConstraints.Add((cons.Coefficients, cons.RightHandSide));
                }
                else if (cons.Operator == ">=")
                {
                    var negatedCoeffs = cons.Coefficients.ToDictionary(k => k.Key, v => -v.Value);
                    processedConstraints.Add((negatedCoeffs, -cons.RightHandSide));
                }
                else if (cons.Operator == "=")
                {
                    processedConstraints.Add((cons.Coefficients, cons.RightHandSide));
                    var negatedCoeffs = cons.Coefficients.ToDictionary(k => k.Key, v => -v.Value);
                    processedConstraints.Add((negatedCoeffs, -cons.RightHandSide));
                }
                else
                {
                    processedConstraints.Add((cons.Coefficients, cons.RightHandSide));
                }
            }
            _numConstraints = processedConstraints.Count;
            int totalCols = _numVars + _numConstraints + 1;
            _tableau = new double[_numConstraints + 1, totalCols];
            for (int i = 0; i < _numConstraints; i++)
            {
                var (coeffs, rhs) = processedConstraints[i];
                for (int j = 0; j < _numVars; j++)
                {
                    string varName = model.Variables[j].Name;
                    if (coeffs.ContainsKey(varName))
                        _tableau[i, j] = coeffs[varName];
                }
                _tableau[i, _numVars + i] = 1.0;
                _tableau[i, totalCols - 1] = rhs;
            }
            bool isMin = string.Equals(
                model.ObjectiveType,
                "Min",
                StringComparison.OrdinalIgnoreCase
            );
            for (int j = 0; j < _numVars; j++)
            {
                double coeff = model.Variables[j].Coefficient;
                if (isMin)
                    coeff = -coeff;
                _tableau[_numConstraints, j] = -coeff;
            }
        }

        private int FindEnteringColumn()
        {
            int colIndex = -1;
            double minVal = 0;
            for (int j = 0; j < _tableau.GetLength(1) - 1; j++)
            {
                if (_tableau[_numConstraints, j] < minVal)
                {
                    minVal = _tableau[_numConstraints, j];
                    colIndex = j;
                }
            }
            return colIndex;
        }

        private int FindLeavingRow(int colIndex)
        {
            int rowIndex = -1;
            double minRatio = double.MaxValue;
            for (int i = 0; i < _numConstraints; i++)
            {
                double coeff = _tableau[i, colIndex];
                double rhs = _tableau[i, _tableau.GetLength(1) - 1];
                if (coeff > 0)
                {
                    double ratio = rhs / coeff;
                    if (ratio < minRatio)
                    {
                        minRatio = ratio;
                        rowIndex = i;
                    }
                }
            }
            return rowIndex;
        }

        private void Pivot(int pivotRow, int pivotCol)
        {
            double pivotValue = _tableau[pivotRow, pivotCol];
            int totalCols = _tableau.GetLength(1);
            for (int j = 0; j < totalCols; j++)
            {
                _tableau[pivotRow, j] /= pivotValue;
            }
            for (int i = 0; i <= _numConstraints; i++)
            {
                if (i != pivotRow)
                {
                    double factor = _tableau[i, pivotCol];
                    for (int j = 0; j < totalCols; j++)
                    {
                        _tableau[i, j] -= factor * _tableau[pivotRow, j];
                    }
                }
            }
        }

        private Solution ExtractSolution(LPQModel model)
        {
            var sol = new Solution
            {
                IsOptimal = true,
                VariableValues = new Dictionary<string, double>(),
                MaxValue = _tableau[_numConstraints, _tableau.GetLength(1) - 1],
                Message = "Success",
            };
            int totalRows = _tableau.GetLength(0);
            int rhsCol = _tableau.GetLength(1) - 1;
            for (int j = 0; j < _numVars; j++)
            {
                string varName = model.Variables[j].Name;
                double value = 0;
                int oneCount = 0;
                int zeroCount = 0;
                int oneRowIndex = -1;
                for (int i = 0; i < totalRows; i++)
                {
                    double val = _tableau[i, j];
                    if (Math.Abs(val - 1.0) < 1e-9)
                    {
                        oneCount++;
                        oneRowIndex = i;
                    }
                    else if (Math.Abs(val) < 1e-9)
                    {
                        zeroCount++;
                    }
                }
                if (
                    oneCount == 1
                    && (oneCount + zeroCount == totalRows)
                    && oneRowIndex < _numConstraints
                )
                {
                    value = _tableau[oneRowIndex, rhsCol];
                }
                sol.VariableValues[varName] = value;
            }
            return sol;
        }
    }
}
