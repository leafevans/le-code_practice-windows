using LPQ.Algorithm;
using LPQ.Domain;
using Microsoft.AspNetCore.Mvc;

namespace LPQ.API.Controllers
{
    [ApiController]
    [Route("api/[controller]")]
    public class SolverController(SimplexSolver solver) : ControllerBase
    {
        private readonly SimplexSolver _solver = solver;

        [HttpPost("solve")]
        public IActionResult Solve([FromBody] LPQModel model)
        {
            if (model == null || model.Variables.Count == 0)
            {
                return BadRequest("模型数据无效");
            }
            Solution result = _solver.Compute(model);
            return Ok(result);
        }
    }
}
