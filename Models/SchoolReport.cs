using System.ComponentModel.DataAnnotations;

namespace prevention_productivity.Models
{
    public class SchoolReport
    {
        [Key]
        public int SchoolReportId { get; set; }
        public string Curriculum { get; set; }
        public string LessonPlan { get; set; }
        public string SchoolName { get; set; }
        public string TopicsCovered { get; set; }
        public string StudentList { get; set; }
        public string Challenges { get; set; }
        public string Successes { get; set; }
        public string Notes { get; set; }
        public ApprovalStatus Status { get; set; }
        public string TeamMemberId { get; set; }
    }
}
