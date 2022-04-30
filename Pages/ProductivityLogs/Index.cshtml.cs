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
        public IList<ApplicationUser> TeamMember { get; set; }
        public IList<GrantProgram> Grants { get; set; }

        public async Task OnGetAsync()
        {
            var grants = from g in _context.GrantProgram
                         select g;
            var productivityLogs = from m in _context.ProductivityLog
                                   select m;
            var teamMembers = from m in _context.Users
                              select m;

            var isAuthorized = User.IsInRole(Constants.AdminRole);

            var currentUserId = UserManager.GetUserId(User);

            if (!isAuthorized)
            {
                productivityLogs = productivityLogs.Where(c => c.TeamMemberID == currentUserId);
            }
            
            ProductivityLog = await productivityLogs.ToListAsync();
            TeamMember = await teamMembers.ToListAsync();
            Grants = await grants.ToListAsync();
        }
    }
}
