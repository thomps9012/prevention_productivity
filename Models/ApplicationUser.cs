using Microsoft.AspNetCore.Identity;

namespace prevention_productivity.Models
{
    public class ApplicationUser : IdentityUser
    {
        public string FirstName { get; set; }
        public string LastName { get; set; }
        
        public string FullName
        {
            get
            {
                return FirstName + " " + LastName;
            }
        }
        
    }
}
