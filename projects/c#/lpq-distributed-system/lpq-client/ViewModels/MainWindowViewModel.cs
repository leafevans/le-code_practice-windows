using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.Net.Http;
using System.Net.Http.Json;
using System.Threading.Tasks;
using CommunityToolkit.Mvvm.ComponentModel;
using CommunityToolkit.Mvvm.Input;

namespace lpq_client.ViewModels
{
    public class CellViewModel : ObservableObject
    {
        private string _value = "0";
        public string Value
        {
            get => _value;
            set { SetProperty(ref _value, value); }
        }
        public bool IsReadOnly { get; set; } = false;
        public string Background { get; set; } = "White";
    }

    public class RowViewModel : ObservableObject
    {
        public string Header { get; set; } = "";
        public ObservableCollection<CellViewModel> Cells { get; set; } = new();
    }

    // DTOs
    public class Variable
    {
        public string Name { get; set; } = "";
        public double Coefficient { get; set; }
    }

    public class Constraint
    {
        public Dictionary<string, double> Coefficients { get; set; } = new();
        public double RightHandSide { get; set; }
    }

    public class LPQModel
    {
        public string Description { get; set; } = "";
        public string ObjectiveType { get; set; } = "Max";
        public List<Variable> Variables { get; set; } = new();
        public List<Constraint> Constraints { get; set; } = new();
    }

    public class Solution
    {
        public bool IsOptimal { get; set; }
        public double MaxValue { get; set; }
        public Dictionary<string, double>? VariableValues { get; set; }
        public string Message { get; set; } = "";
    }

    public partial class MainWindowViewModel : ViewModelBase
    {
        private const string _apiUrl = "http://localhost:5278/api/Solver/solve";

        [ObservableProperty]
        private string _problemDescription = "在此输入问题描述...";

        [ObservableProperty]
        private int _numVars = 2;

        [ObservableProperty]
        private int _numCons = 2;

        [ObservableProperty]
        private bool _isMax = true;

        [ObservableProperty]
        private string _mathModelText = "max z = ...";

        [ObservableProperty]
        private string _resultLog = "";

        [ObservableProperty]
        private bool _showResultPopup = false;

        public ObservableCollection<RowViewModel> GridRows { get; } = new();
        public ObservableCollection<string> ColHeaders { get; } = new();

        public MainWindowViewModel()
        {
            Generate();
        }

        [RelayCommand]
        public void Generate()
        {
            GridRows.Clear();
            ColHeaders.Clear();

            ColHeaders.Add("资源名称");
            for (int j = 0; j < NumVars; j++)
                ColHeaders.Add($"x{j + 1}");
            ColHeaders.Add("限制 (RHS)");

            for (int i = 0; i < NumCons; i++)
            {
                var row = new RowViewModel { Header = $"资源 {Convert.ToChar('A' + i)}" };
                for (int j = 0; j < NumVars; j++)
                    row.Cells.Add(new CellViewModel());
                row.Cells.Add(new CellViewModel { Background = "#E6F7FF" });
                GridRows.Add(row);
            }

            var objRow = new RowViewModel { Header = "单位利润" };
            for (int j = 0; j < NumVars; j++)
                objRow.Cells.Add(new CellViewModel { Background = "#F6FFED" });
            objRow.Cells.Add(
                new CellViewModel
                {
                    Value = "-",
                    IsReadOnly = true,
                    Background = "#EEEEEE",
                }
            );
            GridRows.Add(objRow);

            UpdateMathModelPreview();
        }

        [RelayCommand]
        public void UpdateMathModelPreview()
        {
            string type = IsMax ? "max" : "min";
            string text = $"{type} z = ";

            if (GridRows.Count > 0)
            {
                var objRow = GridRows.Last();
                var parts = new List<string>();
                for (int j = 0; j < NumVars; j++)
                {
                    string val = objRow.Cells[j].Value;
                    if (val != "0")
                        parts.Add($"{val}x{j + 1}");
                }
                text += string.Join(" + ", parts);
            }

            text += "\n\ns.t.\n";

            for (int i = 0; i < NumCons && i < GridRows.Count - 1; i++)
            {
                var row = GridRows[i];
                var parts = new List<string>();
                for (int j = 0; j < NumVars; j++)
                {
                    string val = row.Cells[j].Value;
                    if (val != "0")
                        parts.Add($"{val}x{j + 1}");
                }
                string lhs = parts.Count > 0 ? string.Join(" + ", parts) : "0";
                string rhs = row.Cells[NumVars].Value;
                text += $"   {lhs} <= {rhs}\n";
            }

            text += "   xi >= 0";
            MathModelText = text;
        }

        [RelayCommand]
        public async Task SolveAsync()
        {
            try
            {
                UpdateMathModelPreview();
                var model = BuildModel();

                using var client = new HttpClient();
                var response = await client.PostAsJsonAsync(_apiUrl, model);

                if (response.IsSuccessStatusCode)
                {
                    var result = await response.Content.ReadFromJsonAsync<Solution>();
                    ShowResult(result);
                }
                else
                {
                    ResultLog = $"❌ 服务器错误: {response.StatusCode}";
                    ShowResultPopup = true;
                }
            }
            catch (Exception ex)
            {
                ResultLog = $"❌ 连接失败: {ex.Message}\n请检查后端是否运行";
                ShowResultPopup = true;
            }
        }

        [RelayCommand]
        public void ClosePopup() => ShowResultPopup = false;

        private LPQModel BuildModel()
        {
            var model = new LPQModel
            {
                Description = ProblemDescription,
                ObjectiveType = IsMax ? "Max" : "Min",
            };

            var objRow = GridRows.Last();
            for (int j = 0; j < NumVars; j++)
            {
                double.TryParse(objRow.Cells[j].Value, out double val);
                model.Variables.Add(new Variable { Name = $"x{j + 1}", Coefficient = val });
            }

            for (int i = 0; i < NumCons; i++)
            {
                var row = GridRows[i];
                double.TryParse(row.Cells[NumVars].Value, out double rhs);
                var constraint = new Constraint { RightHandSide = rhs };
                for (int j = 0; j < NumVars; j++)
                {
                    double.TryParse(row.Cells[j].Value, out double val);
                    if (val != 0)
                        constraint.Coefficients.Add($"x{j + 1}", val);
                }
                model.Constraints.Add(constraint);
            }
            return model;
        }

        private void ShowResult(Solution? result)
        {
            if (result == null)
                return;
            string log = $"规划方案结果:\n\n";
            log += $"{(IsMax ? "max" : "min")} z = {result.MaxValue:F2}\n\n";
            if (result.VariableValues != null)
            {
                foreach (var kvp in result.VariableValues)
                {
                    log += $"<{kvp.Key}> = {kvp.Value:F2}\n";
                }
            }
            ResultLog = log;
            ShowResultPopup = true;
        }
    }
}
