using System.Collections.Generic;

namespace LPQ.Domain
{
    public class Variable
    {
        public string Name { get; set; }
        public double Coefficient { get; set; }
    }

    public class Constraint
    {
        public Dictionary<string, double> Coefficients { get; set; }
        public string Operator { get; set; }
        public double RightHandSide { get; set; }
    }

    public class LPQModel
    {
        public string Description { get; set; }
        public string ObjectiveType { get; set; } = "Max";
        public List<Variable> Variables { get; set; }
        public List<Constraint> Constraints { get; set; }
    }

    public class Solution
    {
        public bool IsOptimal { get; set; }
        public double MaxValue { get; set; }
        public Dictionary<string, double> VariableValues { get; set; }
        public string Message { get; set; }
    }
}
