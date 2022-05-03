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
        public string DateSort { get; set; }
        public string TeamMemberSearch { get; set; }
        public string GrantSearch { get; set; }
        public string StatusSearch { get; set; }

        public async Task OnGetAsync(string sortOrder, 
            string teamMemberSearch, 
            string grantSearch,
            string statusSearch)
        {
            DateSort = sortOrder == "asc" ? "date_desc" : "asc";
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
            if (!string.IsNullOrEmpty(teamMemberSearch))
            {
                productivityLogs = productivityLogs.Where(c => c.TeamMemberID == teamMemberSearch);
            }
            if (!string.IsNullOrEmpty(grantSearch))
            {
                productivityLogs = productivityLogs.Where(c => c.GrantProgramID == Int64.Parse(grantSearch));
            }
            if (!string.IsNullOrEmpty(statusSearch))
            {
                switch (statusSearch)
                {
                    case "Approved":
                        productivityLogs = productivityLogs.Where(c => c.Status == ApprovalStatus.Approved);
                        break;
                    case "Pending":
                        productivityLogs = productivityLogs.Where(c => c.Status == ApprovalStatus.Pending);
                        break;
                    case "Rejected":
                        productivityLogs = productivityLogs.Where(c => c.Status == ApprovalStatus.Rejected);
                        break;
                }
            }
            switch (sortOrder)
            {
                case "asc":
                    productivityLogs = productivityLogs.OrderBy(c => c.Date);
                    break;
                case "date_desc":
                    productivityLogs = productivityLogs.OrderByDescending(c => c.Date);
                    break;
                default:
                    productivityLogs = productivityLogs.OrderBy(c => c.Date);
                    break;
            }
            
            ProductivityLog = await productivityLogs.ToListAsync();
            TeamMember = await teamMembers.ToListAsync();
            Grants = await grants.ToListAsync();
        }
    }
}
