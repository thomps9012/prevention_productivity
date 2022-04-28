using System.ComponentModel.DataAnnotations;
using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
using prevention_productivity.Data;

namespace prevention_productivity.Models
{
    public class GrantProgram
    {
        [Key]
        public int Id { get; set; }
        public string AwardNumber { get; set; }
        public string Name { get; set; }
        public string Description { get; set; }
        public DateTime StartDate { get; set; }
        public DateTime EndDate { get; set; }
        public ICollection<ApplicationUser> TeamMembers { get; set; }
        public ICollection<Event> Events { get; set; }
        public ICollection<ProductivityLog> ProductivityLogs { get; set; }

    }
}
