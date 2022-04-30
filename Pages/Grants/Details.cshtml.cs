#nullable disable
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Microsoft.EntityFrameworkCore;
using prevention_productivity.Data;
using prevention_productivity.Models;

namespace prevention_productivity.Pages.Grants
{
    public class DetailsModel : PageModel
    {
        private readonly prevention_productivity.Data.ApplicationDbContext _context;

        public DetailsModel(prevention_productivity.Data.ApplicationDbContext context)
        {
            _context = context;
        }

        public GrantProgram GrantProgram { get; set; }

        public async Task<IActionResult> OnGetAsync(int? id)
        {
            if (id == null)
            {
                return NotFound();
            }

            GrantProgram = await _context.GrantProgram.FirstOrDefaultAsync(m => m.Id == id);

            if (GrantProgram == null)
            {
                return NotFound();
            }
            return Page();
        }
    }
}
