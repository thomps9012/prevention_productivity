using System.ComponentModel.DataAnnotations;

namespace prevention_productivity.Models
{
    public class ProductivityLog
    {
        [Key]
        public int LogID { get; set; }
        [Required]
        public System.DateTime Date { get; set; }
        [Required]
        public string TeamMemberID { get; set; }
        [Required]
        public int GrantProgramID { get; set; }
        [Required]
        public GrantProgram GrantProgram { get; set; }
        public string FocusArea { get; set; }
        [Required]
        public string Actions { get; set; }

        [Required]
        public string Successes { get; set; }
        [Required]
        public string Improvements { get; set; }
        public string NextSteps { get; set; }
        [Required]
        public string Challenges { get; set; }
        [Required]
        public string Notes { get; set; }
        public ApprovalStatus Status { get; set; }
    }
    public enum ApprovalStatus
    {
        Pending,
        Approved,
        Rejected
    }
}
