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

        public async Task<IActionResult> OnGetAsync(int id)
        {

            ProductivityLog? _log = await _context.ProductivityLog.FirstOrDefaultAsync(m => m.LogID == id);

            if (_log == null)
            {
                return NotFound();
            }
            ProductivityLog = _log;
            
            var isAdmin = User.IsInRole(Constants.ProductivityLogsAdminRole);
            
            var currentUserId = UserManager.GetUserId(User);
            
            if (!isAdmin 
                && currentUserId != ProductivityLog.TeamMemberID 
                && ProductivityLog.Status != ApprovalStatus.Approved)
            {
                return Forbid();
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
                return Forbid();
            }
            log.Status = status;
            _context.ProductivityLog.Update(log);
            await _context.SaveChangesAsync();
            return RedirectToPage("./Index");
        }
    }
}
