#nullable disable
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Microsoft.AspNetCore.Mvc.Rendering;
using Microsoft.EntityFrameworkCore;
using prevention_productivity.Data;
using prevention_productivity.Models;

namespace prevention_productivity.Pages.Events.Summary
{
    public class EditModel : PageModel
    {
        private readonly prevention_productivity.Data.ApplicationDbContext _context;

        public EditModel(prevention_productivity.Data.ApplicationDbContext context)
        {
            _context = context;
        }

        [BindProperty]
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
           ViewData["EventId"] = new SelectList(_context.Event, "Id", "Id");
           ViewData["TeamMemberID"] = new SelectList(_context.Users, "Id", "Id");
            return Page();
        }

        // To protect from overposting attacks, enable the specific properties you want to bind to.
        // For more details, see https://aka.ms/RazorPagesCRUD.
        public async Task<IActionResult> OnPostAsync()
        {
            if (!ModelState.IsValid)
            {
                return Page();
            }

            _context.Attach(EventSummary).State = EntityState.Modified;

            try
            {
                await _context.SaveChangesAsync();
            }
            catch (DbUpdateConcurrencyException)
            {
                if (!EventSummaryExists(EventSummary.EventSummaryId))
                {
                    return NotFound();
                }
                else
                {
                    throw;
                }
            }

            return RedirectToPage("./Index");
        }

        private bool EventSummaryExists(int id)
        {
            return _context.EventSummary.Any(e => e.EventSummaryId == id);
        }
    }
}
