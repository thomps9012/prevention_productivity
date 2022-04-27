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

namespace prevention_productivity.Pages.SchoolReports
{
    public class DeleteModel : PageModel
    {
        private readonly prevention_productivity.Data.ApplicationDbContext _context;

        public DeleteModel(prevention_productivity.Data.ApplicationDbContext context)
        {
            _context = context;
        }

        [BindProperty]
        public SchoolReport SchoolReport { get; set; }

        public async Task<IActionResult> OnGetAsync(int? id)
        {
            if (id == null)
            {
                return NotFound();
            }

            SchoolReport = await _context.SchoolReport.FirstOrDefaultAsync(m => m.SchoolReportId == id);

            if (SchoolReport == null)
            {
                return NotFound();
            }
            return Page();
        }

        public async Task<IActionResult> OnPostAsync(int? id)
        {
            if (id == null)
            {
                return NotFound();
            }

            SchoolReport = await _context.SchoolReport.FindAsync(id);

            if (SchoolReport != null)
            {
                _context.SchoolReport.Remove(SchoolReport);
                await _context.SaveChangesAsync();
            }

            return RedirectToPage("./Index");
        }
    }
}
