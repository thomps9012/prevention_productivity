using System.ComponentModel.DataAnnotations;

namespace prevention_productivity.Models
{
    public class EventSummary
    {
        [Key]
        public int EventSummaryId { get; set; }
        public int EventId { get; set; }
        public Event Event { get; set; }
        public int Attendees { get; set; }
        public string Challenges { get; set; }
        public string Outcomes { get; set; }
        public string Notes { get; set; }
        public string NextSteps { get; set; }
        public string TeamMemberID { get; set; }
        public ApplicationUser TeamMember { get; set; }
        public ApprovalStatus Status { get; set; }
    }
}
