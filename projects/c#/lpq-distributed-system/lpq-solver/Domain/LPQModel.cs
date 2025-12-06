using System.Collections.Generic;

namespace LPQ.Domain
{
    public class Variable
    {
        public string Name { get; set; } = string.Empty;
        public double Coefficient { get; set; }
    }

    public class Constraint
    {
        public Dictionary<string, double> Coefficients { get; set; } =
            new Dictionary<string, double>();
        public string Operator { get; set; } = "<=";
        public double RightHandSide { get; set; }
    }

    public class LPQModel
    {
        public string Description { get; set; } = string.Empty;
        public string ObjectiveType { get; set; } = "Max";

        // 初始化列表
        public List<Variable> Variables { get; set; } = new List<Variable>();
        public List<Constraint> Constraints { get; set; } = new List<Constraint>();
    }

    public class Solution
    {
        public bool IsOptimal { get; set; }
        public double MaxValue { get; set; }

        public Dictionary<string, double>? VariableValues { get; set; }
        public string Message { get; set; } = string.Empty;
    }
}
