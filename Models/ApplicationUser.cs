using Microsoft.AspNetCore.Identity;

namespace prevention_productivity.Models
{
    public class ApplicationUser : IdentityUser
    {
        public string FirstName { get; set; }
        public string LastName { get; set; }
        
    }
}
