using System.ComponentModel.DataAnnotations;

namespace prevention_productivity.Models
{
    public class Comment
    {
        [Key]
        public int Id { get; set; }
        public string AuthorId { get; set; }
        public string Content { get; set; }
        public DateTime CreatedAt { get; set; }
        public string ItemId { get; set; }
    }
}
