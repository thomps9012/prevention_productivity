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

namespace prevention_productivity.Pages.Events.Summary
{
    public class DetailsModel : PageModel
    {
        private readonly ApplicationDbContext _context;

        public DetailsModel(ApplicationDbContext context)
        {
            _context = context;
        }

        public EventSummary EventSummary { get; set; }

        public async Task<IActionResult> OnGetAsync(int? id)
        {
            if (id == null)
            {
                return NotFound();
            }

            EventSummary = await _context.EventSummary
                .Include(e => e.Event)
                .Include(e => e.TeamMember).FirstOrDefaultAsync(m => m.EventSummaryId == id);

            if (EventSummary == null)
            {
                return NotFound();
            }
            return Page();
        }

        public async Task<IActionResult> OnPostAsync(int? id)
        {
            if(id == null){
                return NotFound();
            }
            EventSummary = await _context.EventSummary.FindAsync(id);
            if (EventSummary != null)
            {
                _context.EventSummary.Remove(EventSummary);
                await _context.SaveChangesAsync();
            }
            return RedirectToPage("./Index");
        }
    }
}
