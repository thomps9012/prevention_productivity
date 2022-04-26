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

namespace prevention_productivity.Pages.Events
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

        public Event Event { get; set; }
        public EventSummary EventSummary { get; set; }

        public async Task<IActionResult> OnGetAsync(int? id)
        {
            if (id == null)
            {
                return NotFound();
            }

            Event = await _context.Event
                .Include(a => a.GrantProgram).FirstOrDefaultAsync(m => m.Id == id);

            EventSummary = await _context.EventSummary.Include(a => a.Event).FirstOrDefaultAsync(m => m.EventId == id);

            var isAdmin = User.IsInRole(Constants.AdminRole);
            var currentUserId = UserManager.GetUserId(User);

            if (Event == null)
            {
                return NotFound();
            }
            if (!isAdmin
                && currentUserId != Event.EventLead
                && Event.Status != ApprovalStatus.Approved)
            {
                return Forbid();
            }
            return Page();
        }
        public async Task<IActionResult> OnPostAsync(int id, ApprovalStatus status)
        {
            var _event = await _context.Event.FirstOrDefaultAsync(m => m.Id == id);
            if(_event == null)
            {
            return NotFound();
        }
        var operation = (status == ApprovalStatus.Approved) 
            ? AuthOperations.Approve 
            : AuthOperations.Reject;
        
        var isAuthorized = await AuthorizationService.AuthorizeAsync(User, _event, operation);
        if(!isAuthorized.Succeeded)
        {
            return Forbid();
    }
    _event.Status = status;
    _context.Event.Update(_event);
    await _context.SaveChangesAsync();
    return RedirectToPage("./Index");
    }
}
}
