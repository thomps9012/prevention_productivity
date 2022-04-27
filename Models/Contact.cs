using System.ComponentModel.DataAnnotations;

namespace prevention_productivity.Models
{
    public class Contact
    {
        [Key]
        public int ContactId { get; set; }
        public ContactType Type { get; set; }
        public string FirstName { get; set; }
        public string LastName { get; set; }
        public string FullName
        {
            get
            {
                return FirstName + " " + LastName;
            }
        }
        [DataType(DataType.EmailAddress)]
        public string Email { get; set; }
        [DataType(DataType.PhoneNumber)]
        public string Phone { get; set; }
        public string AffiliatedOrg { get; set; }
        public DateTime CreatedAt {
            get
            {
                return DateTime.Now;
            }
        }
    }
        public enum ContactType
        {
            Student,
            Parent,
            Teacher,
            NonProfit,
            Public,
            Private,
            Other
        }
}
