#nullable disable
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
using prevention_productivity.Authorization;
using prevention_productivity.Data;
using prevention_productivity.Models;

namespace prevention_productivity.Pages.ProductivityLogs
{
    public class IndexModel : DI_BasePageModel
    {
        private readonly ApplicationDbContext _context;

        public IndexModel(ApplicationDbContext context,
            IAuthorizationService authorizationService,
            UserManager<ApplicationUser> userManager)
            : base(context, authorizationService, userManager)
        {
            _context = context;
        }
       

        public IList<ProductivityLog> ProductivityLog { get;set; }

        public async Task OnGetAsync()
        {
            var productivityLogs = from m in _context.ProductivityLog
                                   select m;
            
            var isAuthorized = User.IsInRole(Constants.ProductivityLogsAdminRole);

            var currentUserId = UserManager.GetUserId(User);

            if (!isAuthorized)
            {
                productivityLogs = productivityLogs.Where(c => c.Status == ApprovalStatus.Approved
                                                || c.TeamMemberID == currentUserId);
            }
            
            ProductivityLog = await productivityLogs.ToListAsync();
        }
    }
}
