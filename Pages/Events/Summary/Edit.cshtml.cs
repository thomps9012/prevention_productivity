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

                return Page();
            } else
            {
                return Forbid();
            }
        }

        public async Task<IActionResult> OnPostAsync(string action)
        {
           // if (action == "delete") {

//                var isAuthorized = await AuthorizationService.AuthorizeAsync(
  //                                                       User, EventSummary,
    //                                                     AuthOperations.Delete);
      //          if (!isAuthorized.Succeeded)
        //        {
          //          return Forbid();
            //    }
            //

              //  Context.EventSummary.Remove(EventSummary);
                //await Context.SaveChangesAsync();


               // return RedirectToPage("../Index");
           // }
           // else if (action == "save")
           // {

                //   if (!ModelState.IsValid)
                //  {
                //     return Page();
                // }

                var summaryToUpdate = await Context.EventSummary.AsNoTracking().FirstOrDefaultAsync(s => s.EventSummaryId == EventSummary.EventSummaryId);
                if (summaryToUpdate == null)
                {
                    return NotFound();
                }
                var isAuthorized = await AuthorizationService.AuthorizeAsync(
                                                     User, summaryToUpdate,
                                                     AuthOperations.Update);
                if (!isAuthorized.Succeeded)
                {
                    return Forbid();
                }
                Context.Attach(EventSummary).State = EntityState.Modified;
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
                await Context.SaveChangesAsync();
                return RedirectToPage("./Index");
           // }
         //   return Page();
        }
    }
}