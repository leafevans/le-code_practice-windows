namespace LPQ.API.Controllers
{
    [ApiController]
    [Route("api/[controller]")]
    public class SolverController : ControllerBase
    {
        [HttpPost("solve")]
        public IActionResult Solve([FromBody] LPQModel model)
        {
            if (model == null || model.Variables == null || model.Variables.Count == 0)
            {
                return BadRequest("模型数据无效");
            }

            try
            {
                var solver = new SimplexSolver();
                Solution result = solver.Compute(model);
                return Ok(result);
            }
            catch (System.Exception ex)
            {
                return StatusCode(500, "内部计算错误: " + ex.Message);
            }
        }
    }
}
