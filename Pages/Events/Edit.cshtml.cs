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

namespace prevention_productivity.Pages.Events
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
        public Event Event { get; set; }
        public IList<GrantProgram> Grants { get; set; }

        public async Task<IActionResult> OnGetAsync(int? id)
        {
            if (id == null)
            {
                return NotFound();
            }

            Event = await _context.Event
                .Include(a => a.GrantProgram).FirstOrDefaultAsync(m => m.Id == id);
            Grants = await _context.GrantProgram.ToListAsync();
            if (Event == null)
            {
                return NotFound();
            }
          
            if ((await AuthorizationService.AuthorizeAsync(User, Event, AuthOperations.Update)).Succeeded)
            {
                return Page();
            }
            else
            {
                return Forbid();
            }
        }

        // To protect from overposting attacks, enable the specific properties you want to bind to.
        // For more details, see https://aka.ms/RazorPagesCRUD.
        public async Task<IActionResult> OnPostAsync()
        {
            //  if (!ModelState.IsValid)
            // {
            //    return Page();
            // }
            var eventToUpdate = await Context.Event.AsNoTracking().FirstOrDefaultAsync(m => m.Id == Event.Id);
        if (eventToUpdate == null)
            {
            return NotFound();
        }
        var isAuthorized = await AuthorizationService.AuthorizeAsync(User, Event, AuthOperations.Update);
            if (!isAuthorized.Succeeded)
            {
                return Forbid();
            }

            Context.Attach(Event).State = EntityState.Modified;
            if (Event.Status == ApprovalStatus.Approved)
            {
                var canApprove = await AuthorizationService.AuthorizeAsync(
                    User,
                    Event,
                    AuthOperations.Approve);
                if (!canApprove.Succeeded)
                {
                    Event.Status = ApprovalStatus.Pending;
                }
            }
            await Context.SaveChangesAsync();
            return RedirectToPage("./Details", new { id = Event.Id });
        }
    }
}