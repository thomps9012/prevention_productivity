#nullable disable
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Microsoft.EntityFrameworkCore;
using prevention_productivity.Authorization;
using prevention_productivity.Data;
using prevention_productivity.Models;
using prevention_productivity.Pages.ProductivityLogs;

namespace prevention_productivity.Pages.Events.Summary
{
    public class DetailsModel : DI_BasePageModel
    {
        private readonly ApplicationDbContext _context;

        public DetailsModel(ApplicationDbContext context,
             IAuthorizationService authorizationService,
             UserManager<ApplicationUser> userManager)
             : base(context, authorizationService, userManager)
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
            var isAdmin = User.IsInRole(Constants.AdminRole);
            var currentUserId = UserManager.GetUserId(User);
            if (EventSummary == null)
            {
                return NotFound();
            }
            if (!isAdmin
                && currentUserId != EventSummary.TeamMemberID
                && EventSummary.Status != ApprovalStatus.Approved)
            {
                return Forbid();
            }            
            return Page();
        }

        public async Task<IActionResult> OnPostAsync(int? id, ApprovalStatus status)
        {
            if(id == null){
                return NotFound();
            }
            EventSummary = await _context.EventSummary.FindAsync(id);
      
            if (EventSummary == null)
            {
                return NotFound();
            }
            var operation = (status == ApprovalStatus.Approved)
            ? AuthOperations.Approve
            : AuthOperations.Reject;

            var isAuthorized = await AuthorizationService.AuthorizeAsync(User, EventSummary, operation);
            if (!isAuthorized.Succeeded)
            {
                return Forbid();
            }


            EventSummary.Status = status;
            _context.EventSummary.Update(EventSummary);
            await _context.SaveChangesAsync();
            return RedirectToPage("../Index");
        }
    }
}
