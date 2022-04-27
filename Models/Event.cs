using System.ComponentModel.DataAnnotations;

namespace prevention_productivity.Models
{
    public class Event
    {
        [Key]
        public int Id { get; set; }
        public string Name { get; set; }
        public string Description { get; set; }
        public string Location { get; set; }
        public string LocationDetail { get; set; }
        public int GrantProgramId { get; set; }
        public GrantProgram GrantProgram { get; set; }
        public bool PublicEvent { get; set; }
        public bool RSVPRequired { get; set; }
        public bool AnnualEvent { get; set; }
        public bool NewEvent { get; set; }
        public DateTime EventStart { get; set; }
        public double Duration { get; set; }
        public DateTime SetUpTime { get; set; }
        public DateTime CleanUpTime { get; set; }
        public string Agenda { get; set; }
        public string TargetPopulation { get; set; }
        public string AgeGroup { get; set; }
        public string PartingGift { get; set; }
        public string Raffle { get; set; }
        public string Marketing { get; set; }
        public string SpecialOrderGear { get; set; }
        public string SpecialOrderDesignElements { get; set; }
        public string Performance { get; set; }
        public string Vendors { get; set; }
        public string FoodBeverage { get; set; }
        public string Caterer { get; set; }
        public int FoodHeadCount { get; set; }
        public string OtherSpecialOrder { get; set; }
        public string PreventionTeamMembers { get; set; }
        public double StaffTimeCommitment { get; set; }
        public string StaffDuties { get; set; }
        public string NORAClients { get; set; }
        public bool Volunteers { get; set; }
        public string VolunteerList { get; set; }
        public string Supplies { get; set; }
        public double EventBudget { get; set; }
        public string EventAffiliates { get; set; }
        public string Notes { get; set; }
        public string EventLead { get; set; }
        public ApprovalStatus Status { get; set; }
    }
}
