#nullable disable
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Microsoft.AspNetCore.Mvc.Rendering;
using Microsoft.EntityFrameworkCore;
using prevention_productivity.Authorization;
using prevention_productivity.Data;
using prevention_productivity.Models;
using prevention_productivity.Pages.ProductivityLogs;

namespace prevention_productivity.Pages.Events.Summary
{
    public class EditModel : DI_BasePageModel
    {
        private readonly ApplicationDbContext _context;

        public EditModel(ApplicationDbContext context,
            IAuthorizationService authorizationService,
            UserManager<ApplicationUser> userManager)
            : base(context, authorizationService, userManager)
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
            if ((await AuthorizationService.AuthorizeAsync(User, EventSummary, AuthOperations.Update)).Succeeded)
            {
                ViewData["EventId"] = new SelectList(_context.Event, "Id", "Name");
                ViewData["TeamMemberID"] = new SelectList(_context.Users, "Id", "FullName");
                return Page();
            } else
            {
                return Forbid();
            }
        }

        public async Task<IActionResult> OnPostAsync(string action)
        {
            if (action == "delete") { 
           
            var isAuthorized = await AuthorizationService.AuthorizeAsync(
                                                     User, EventSummary,
                                                     AuthOperations.Delete);
            if (!isAuthorized.Succeeded)
            {
                return Forbid();
            }
            
                
                _context.EventSummary.Remove(EventSummary);
                await _context.SaveChangesAsync();
            

            return RedirectToPage("./Index");
            }else
            {
                if (!ModelState.IsValid)
                {
                    return Page();
                }


                if ((await AuthorizationService.AuthorizeAsync(User, EventSummary, AuthOperations.Update)).Succeeded)
                {

                    _context.Attach(EventSummary).State = EntityState.Modified;
                    if (EventSummary.Status == ApprovalStatus.Approved)
                    {
                        var canApprove = await AuthorizationService.AuthorizeAsync(
                            User,
                            EventSummary,
                            AuthOperations.Approve);
                        if (!canApprove.Succeeded)
                        {
                            EventSummary.Status = ApprovalStatus.Pending;
                        }
                    }
                }
                else
                {
                    return Forbid();
                }


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
        }

        private bool EventSummaryExists(int id)
        {
            return _context.EventSummary.Any(e => e.EventSummaryId == id);
        }
    }
}
