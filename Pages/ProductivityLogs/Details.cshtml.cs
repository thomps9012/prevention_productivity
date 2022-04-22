using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using prevention_productivity.Data;
using prevention_productivity.Models;
using prevention_productivity.Authorization;

namespace prevention_productivity.Pages.ProductivityLogs
{
    public class DetailsModel : DI_BasePageModel
    {
        private readonly ApplicationDbContext _context;

        public DetailsModel(ApplicationDbContext context,
            IAuthorizationService authorizationService,
            UserManager<IdentityUser> userManager)
            : base(context, authorizationService, userManager)
        {
            _context = context;
        }

        public ProductivityLog ProductivityLog { get; set; }

        public async Task<IActionResult> OnGetAsync(int? id)
        {
            if (id == null)
            {
                return NotFound();
            }

            ProductivityLog = await _context.ProductivityLog.FirstOrDefaultAsync(m => m.LogID == id);

            if (ProductivityLog == null)
            {
                return NotFound();
            }
            var isAdmin = User.IsInRole(Constants.ProductivityLogsAdminRole);
            var currentUserId = UserManager.GetUserId(User);
            if (!isAdmin 
                && currentUserId != ProductivityLog.TeamMemberID 
                && ProductivityLog.Status != ApprovalStatus.Approved)
            {
                return new ForbidResult();
            }
            return Page();
        }
        public async Task<IActionResult> OnPostAsync(int id, ApprovalStatus status)
        {
            var log = await _context.ProductivityLog.FirstOrDefaultAsync(m => m.LogID == id);
            if (log == null)
            {
                return NotFound();
            }
            var operation = (status == ApprovalStatus.Approved) 
                ? ProductivityLogOperations.Approve 
                : ProductivityLogOperations.Reject;

            var isAuthorized = await AuthorizationService.AuthorizeAsync(User, log, operation);
            if (!isAuthorized.Succeeded)
            {
                return new ForbidResult();
            }
            log.Status = status;
            Context.ProductivityLog.Update(log);
            await Context.SaveChangesAsync();
            return RedirectToPage("./Index");
        }
    }
}
