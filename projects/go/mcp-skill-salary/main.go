package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ====================== 资源与常量 ======================

const taxRule2026 = `2026年个人所得税新规：
1. 个税起征点：5000元/月
2. 应纳税所得额 = 税前工资 - 五险一金 - 5000
3. 税率表（按应纳税所得额）：
   - 0 ~ 3000元：3%
   - 3000 ~ 12000元：10%，速算扣除数 210
   - 12000 ~ 25000元：20%，速算扣除数 1410
   - 25000 ~ 35000元：25%，速算扣除数 2660
4. 五险一金扣除比例：月工资的 10.5%`

const slipTemplate = `工资单
=====
员工姓名：%s
核算月份：%s
税前工资：%.0f 元
考勤扣款：%.0f 元
五险一金：%.2f 元
个税扣除：%.2f 元
实发工资：%.2f 元`

// ====================== 结构体定义 ======================

// SalaryCalcInput 工资核算入参
type SalaryCalcInput struct {
	EmployeeName string `json:"employee_name" jsonschema:"员工姓名，必填"`
	Month        string `json:"month" jsonschema:"核算月份，格式 YYYY-MM，必填"`
	BaseSalary   string `json:"base_salary" jsonschema:"税前月基本工资（元），必填"`
}

// SalaryCalcOutput 工资核算出参
type SalaryCalcOutput struct {
	EmployeeName     string  `json:"employee_name" jsonschema:"员工姓名"`
	Month            string  `json:"month" jsonschema:"核算月份"`
	BaseSalary       float64 `json:"base_salary" jsonschema:"税前工资"`
	AttendanceDeduct float64 `json:"attendance_deduct" jsonschema:"考勤扣款"`
	SocialSecurity   float64 `json:"social_security" jsonschema:"五险一金扣除"`
	Tax              float64 `json:"tax" jsonschema:"个税扣除"`
	FinalSalary      float64 `json:"final_salary" jsonschema:"实发工资"`
	Slip             string  `json:"slip" jsonschema:"工资单文本"`
}

type attendanceInput struct {
	EmployeeName string `json:"employee_name" jsonschema:"员工姓名"`
	Month        string `json:"month" jsonschema:"核算月份"`
}

type taxInput struct {
	BaseSalary float64 `json:"base_salary" jsonschema:"税前工资（已扣除考勤后的金额）"`
}

type attendanceOutput struct {
	Deduct float64 `json:"deduct" jsonschema:"考勤扣款金额"`
}

type taxOutput struct {
	Tax float64 `json:"tax" jsonschema:"个税金额"`
}

// ====================== 原子 Tool：考勤查询 ======================

func queryAttendance(ctx context.Context, req *mcp.CallToolRequest, in attendanceInput) (
	*mcp.CallToolResult, attendanceOutput, error,
) {
	// 模拟考勤系统数据，实际项目中应替换为数据库查询
	out := attendanceOutput{
		Deduct: map[string]float64{
			"张三": 200,
			"李四": 0,
			"王五": 500,
		}[in.EmployeeName],
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("%.2f", out.Deduct)},
		},
	}, out, nil
}

// ====================== 原子 Tool：个税计算 ======================

func calculateTax(ctx context.Context, req *mcp.CallToolRequest, in taxInput) (
	*mcp.CallToolResult, taxOutput, error,
) {
	// 应纳税所得额 = 税前工资 - 五险一金 - 起征点
	taxable := in.BaseSalary - in.BaseSalary*0.105 - 5000
	out := taxOutput{Tax: calcTaxFromTaxable(taxable)}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("%.2f", out.Tax)},
		},
	}, out, nil
}

// calcTaxFromTaxable 根据应纳税所得额计算个税
func calcTaxFromTaxable(taxable float64) float64 {
	if taxable <= 0 {
		return 0
	}
	switch {
	case taxable <= 3000:
		return taxable * 0.03
	case taxable <= 12000:
		return taxable*0.1 - 210
	case taxable <= 25000:
		return taxable*0.2 - 1410
	default:
		return taxable*0.25 - 2660
	}
}

// ====================== Skill：工资核算主流程 ======================

func salaryCalculationSkill(ctx context.Context, req *mcp.CallToolRequest, in SalaryCalcInput) (
	*mcp.CallToolResult, SalaryCalcOutput, error,
) {
	// ----- 步骤 1：入参校验 -----
	if in.EmployeeName == "" {
		return nil, SalaryCalcOutput{}, fmt.Errorf("员工姓名不能为空")
	}
	if in.Month == "" {
		return nil, SalaryCalcOutput{}, fmt.Errorf("核算月份不能为空")
	}
	base, err := strconv.ParseFloat(in.BaseSalary, 64)
	if err != nil || base <= 0 {
		return nil, SalaryCalcOutput{}, fmt.Errorf("税前基本工资必须是大于 0 的数字")
	}

	// ----- 步骤 2：按固化流程执行核算 -----

	// 2.1 考勤扣款
	attendance := map[string]float64{
		"张三": 200,
		"李四": 0,
		"王五": 500,
	}[in.EmployeeName]

	// 2.2 五险一金
	social := base * 0.105

	// 2.3 个税（应纳税所得额 = 税前 - 考勤 - 五险一金 - 起征点）
	taxable := base - attendance - (base-attendance)*0.105 - 5000
	tax := calcTaxFromTaxable(taxable)

	// 2.4 实发工资
	final := base - attendance - social - tax

	// ----- 步骤 3：结果校验 -----
	if final < 0 {
		return nil, SalaryCalcOutput{}, fmt.Errorf("工资核算异常，实发工资不能为负数，请检查输入参数")
	}

	// ----- 步骤 4：生成工资单 -----
	slip := fmt.Sprintf(slipTemplate, in.EmployeeName, in.Month, base, attendance, social, tax, final)

	// ----- 步骤 5：审计日志 -----
	fmt.Fprintf(os.Stderr, "[审计] 员工：%s | 月份：%s | 税前：%.0f | 实发：%.2f\n",
		in.EmployeeName, in.Month, base, final)

	output := SalaryCalcOutput{
		EmployeeName:     in.EmployeeName,
		Month:            in.Month,
		BaseSalary:       base,
		AttendanceDeduct: attendance,
		SocialSecurity:   social,
		Tax:              tax,
		FinalSalary:      final,
		Slip:             slip,
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: slip},
		},
	}, output, nil
}

// ====================== Resource：个税规则 ======================

func taxRuleResource(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{
			{URI: req.Params.URI, MIMEType: "text/plain", Text: taxRule2026},
		},
	}, nil
}

// ====================== 主函数 ======================

func main() {
	server := mcp.NewServer(
		&mcp.Implementation{Name: "salary-calculation-skill", Version: "1.0.0"},
		&mcp.ServerOptions{
			Instructions: "员工月度工资核算服务，根据员工姓名、核算月份、税前基本工资，自动完成考勤扣款、五险一金、个税计算，生成标准工资单。",
			Capabilities: &mcp.ServerCapabilities{
				Resources: &mcp.ResourceCapabilities{Subscribe: false},
				Tools:     &mcp.ToolCapabilities{ListChanged: false},
			},
		},
	)

	// 注册原子 Tool：考勤查询
	mcp.AddTool(server, &mcp.Tool{
		Name:        "query_attendance",
		Description: "查询指定员工在指定月份的考勤扣款金额",
	}, queryAttendance)

	// 注册原子 Tool：个税计算
	mcp.AddTool(server, &mcp.Tool{
		Name:        "calculate_tax",
		Description: "根据税前工资和 2026 年个税新规计算个人所得税",
	}, calculateTax)

	// 注册 Skill：工资核算主流程
	mcp.AddTool(server, &mcp.Tool{
		Name: "salary_calculation",
		Description: `员工月度工资核算（Skill），根据员工姓名、核算月份、税前基本工资，` +
			`按固化流程自动完成：考勤扣款 → 五险一金 → 个税计算 → 生成工资单。` +
			`仅用于员工月度工资核算场景。`,
	}, salaryCalculationSkill)

	// 注册个税规则 Resource
	server.AddResource(&mcp.Resource{
		Name:        "2026年个人所得税新规",
		URI:         "https://company.com/tax/2026",
		Description: "2026 年个人所得税税率表及计算规则",
		MIMEType:    "text/plain",
	}, taxRuleResource)

	// 启动 STDIO 服务
	fmt.Fprintln(os.Stderr, "工资核算 Skill 已启动（STDIO 模式）")
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("服务启动失败：%v", err)
	}
}
