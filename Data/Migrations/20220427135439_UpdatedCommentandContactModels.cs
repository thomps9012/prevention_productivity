using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace prevention_productivity.Data.Migrations
{
    public partial class UpdatedCommentandContactModels : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_Comment_Users_AuthorId",
                table: "Comment");

            migrationBuilder.DropIndex(
                name: "IX_Comment_AuthorId",
                table: "Comment");
            migrationBuilder.CreateTable(
               name: "Contact",
               columns: table => new
               {
                   ContactId = table.Column<int>(type: "int", nullable: false)
                       .Annotation("SqlServer:Identity", "1, 1"),
                   FirstName = table.Column<string>(type: "nvarchar(max)", nullable: false),
                   LastName = table.Column<string>(type: "nvarchar(max)", nullable: false),
                   Email = table.Column<string>(type: "nvarchar(max)", nullable: false),
                   Phone = table.Column<string>(type: "nvarchar(max)", nullable: false),
                   AffiliatedOrg = table.Column<string>(type: "nvarchar(max)", nullable: false),
                   Type = table.Column<int>(type: "int", nullable: false, defaultValue: 0),
               },
               constraints: table =>
               {
                   table.PrimaryKey("PK_Contact", x => x.ContactId);
               });

            migrationBuilder.AlterColumn<string>(
                name: "AuthorId",
                table: "Comment",
                type: "nvarchar(max)",
                nullable: false,
                defaultValue: "",
                oldClrType: typeof(string),
                oldType: "nvarchar(450)",
                oldNullable: true);
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {

            migrationBuilder.AlterColumn<string>(
                name: "AuthorId",
                table: "Comment",
                type: "nvarchar(450)",
                nullable: true,
                oldClrType: typeof(string),
                oldType: "nvarchar(max)");

            migrationBuilder.CreateIndex(
                name: "IX_Comment_AuthorId",
                table: "Comment",
                column: "AuthorId");

            migrationBuilder.AddForeignKey(
                name: "FK_Comment_Users_AuthorId",
                table: "Comment",
                column: "AuthorId",
                principalTable: "Users",
                principalColumn: "Id");
        }
    }
}
