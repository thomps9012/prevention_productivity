

// dotnet aspnet-codegenerator razorpage -m Contact -dc ApplicationDbContext -udl -outDir Pages\Contacts --referenceScriptLibraries

using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
using prevention_productivity.Authorization;
using prevention_productivity.Data;
using prevention_productivity.Models;

namespace prevention_productivity.Data
{
    public static class SeedData
    {
        public static async Task Initialize(IServiceProvider serviceProvider, string testUserPw)
        {
            using (var context = new ApplicationDbContext(
                serviceProvider.GetRequiredService<DbContextOptions<ApplicationDbContext>>()))
            {
                var adminID = await EnsureUser(serviceProvider, testUserPw, "admin@test.com");
                await EnsureRole(serviceProvider, adminID, Constants.ProductivityLogsAdminRole);

                var teamMember1 = await EnsureUser(serviceProvider, testUserPw, "test@test.com");
                await EnsureRole(serviceProvider, teamMember1, Constants.ProductivityLogsUserRole);
                var teamMember2 = await EnsureUser(serviceProvider, testUserPw, "test2@test.com");
                await EnsureRole(serviceProvider, teamMember2, Constants.ProductivityLogsUserRole);
                var teamMember3 = await EnsureUser(serviceProvider, testUserPw, "test3@test.com");
                await EnsureRole(serviceProvider, teamMember3, Constants.ProductivityLogsUserRole);

                SeedDB(context, adminID, teamMember1, teamMember2, teamMember3);
            }
        }

        private static async Task<string> EnsureUser(IServiceProvider serviceProvider,
                                                    string testUserPw, string UserName)
        {
            var userManager = serviceProvider.GetService<UserManager<IdentityUser>>();

            var user = await userManager.FindByNameAsync(UserName);
            if (user == null)
            {
                user = new IdentityUser { UserName = UserName };
                await userManager.CreateAsync(user, testUserPw);
            }
            if (user == null)
            {
                throw new Exception("The password is probably not strong enough!");
            }

            return user.Id;
        }

        private static async Task<IdentityResult> EnsureRole(IServiceProvider serviceProvider,
                                                                      string uid, string role)
        {
            var roleManager = serviceProvider.GetService<RoleManager<IdentityRole>>();
            if (roleManager == null)
            {
                throw new Exception("roleManager null");
            }
            IdentityResult IR = null;
            if (!await roleManager.RoleExistsAsync(role))
            {
                IR = await roleManager.CreateAsync(new IdentityRole(role));
            }

            var userManager = serviceProvider.GetService<UserManager<IdentityUser>>();

            var user = await userManager.FindByIdAsync(uid);
            if (user == null)
            {
                throw new Exception("The testUserPw password was probably not strong enough!");
            }

            IR = await userManager.AddToRoleAsync(user, role);

            return IR;
        }


        public static void SeedDB(ApplicationDbContext context, string adminID, string teamMember1, string teamMember2, string teamMember3)
        {
            if (context.ProductivityLog.Any())
            {
                return;   // DB has been seeded
            }

            context.ProductivityLog.AddRange(
                new ProductivityLog
                {
                    LogID = 1,
                    TeamMemberID = adminID,
                    Date = System.DateTime.Now,
                    FocusArea = "Focus Area 1",
                    Actions = "Actions 1",
                    Successes = "Successes 1",
                    Improvements = "Improvements 1",
                    Challenges = "Challenges 1",
                    NextSteps = "Next Steps 1",
                    Notes = "Notes 1",
                    Status = (ApprovalStatus)2
                },
                new ProductivityLog
                {
                    LogID = 2,
                    TeamMemberID = teamMember1,
                    Date = System.DateTime.Now,
                    FocusArea = "Focus Area 2",
                    Actions = "Actions 2",
                    Successes = "Successes 2",
                    Improvements = "Improvements 2",
                    Challenges = "Challenges 2",
                    NextSteps = "Next Steps 2",
                    Notes = "Notes 2",
                    Status = (ApprovalStatus)0
                },
                new ProductivityLog
                {
                    LogID = 3,
                    TeamMemberID = teamMember2,
                    Date = System.DateTime.Now,
                    FocusArea = "Focus Area 3",
                    Actions = "Actions 3",
                    Successes = "Successes 3",
                    Improvements = "Improvements 3",
                    Challenges = "Challenges 3",
                    NextSteps = "Next Steps 3",
                    Notes = "Notes 3",
                    Status = (ApprovalStatus)1
                },
                new ProductivityLog
                {
                    LogID = 4,
                    TeamMemberID = teamMember3,
                    Date = System.DateTime.Now,
                    FocusArea = "Focus Area 4",
                    Actions = "Actions 4",
                    Successes = "Successes 4",
                    Improvements = "Improvements 4",
                    Challenges = "Challenges 4",
                    NextSteps = "Next Steps 4",
                    Notes = "Notes 4",
                    Status = (ApprovalStatus)0
                },
                new ProductivityLog
                {
                    LogID = 5,
                    TeamMemberID = teamMember1,
                    Date = System.DateTime.Now,
                    FocusArea = "Focus Area 5",
                    Actions = "Actions 5",
                    Successes = "Successes 5",
                    Improvements = "Improvements 5",
                    Challenges = "Challenges 5",
                    NextSteps = "Next Steps 5",
                    Notes = "Notes 5",
                    Status = (ApprovalStatus)1
                },
                new ProductivityLog
                {
                    LogID = 6,
                    TeamMemberID = teamMember2,
                    Date = System.DateTime.Now,
                    FocusArea = "Focus Area 6",
                    Actions = "Actions 6",
                    Successes = "Successes 6",
                    Improvements = "Improvements 6",
                    Challenges = "Challenges 6",
                    NextSteps = "Next Steps 6",
                    Notes = "Notes 6",
                    Status = (ApprovalStatus)0
                },
                new ProductivityLog
                {
                    LogID = 7,
                    TeamMemberID = teamMember3,
                    Date = System.DateTime.Now,
                    FocusArea = "Focus Area 7",
                    Actions = "Actions 7",
                    Successes = "Successes 7",
                    Improvements = "Improvements 7",
                    Challenges = "Challenges 7",
                    NextSteps = "Next Steps 7",
                    Notes = "Notes 7",
                    Status = (ApprovalStatus)1
                }


             );
            context.SaveChanges();
        }

    }
}