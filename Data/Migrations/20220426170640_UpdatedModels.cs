using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace prevention_productivity.Data.Migrations
{
    public partial class UpdatedModels : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_ProductivityLog_GrantProgram_GrantProgramId",
                table: "ProductivityLog");

            migrationBuilder.RenameColumn(
                name: "GrantProgramId",
                table: "ProductivityLog",
                newName: "GrantProgramID");

            migrationBuilder.RenameIndex(
                name: "IX_ProductivityLog_GrantProgramId",
                table: "ProductivityLog",
                newName: "IX_ProductivityLog_GrantProgramID");

            migrationBuilder.AlterColumn<int>(
                name: "GrantProgramID",
                table: "ProductivityLog",
                type: "int",
                nullable: false,
                defaultValue: 0,
                oldClrType: typeof(int),
                oldType: "int",
                oldNullable: true);

            migrationBuilder.CreateTable(
                name: "Conttact",
                columns: table => new
                {
                    ContactId = table.Column<int>(type: "int", nullable: false)
                        .Annotation("SqlServer:Identity", "1, 1"),
                    ContactType = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    FirstName = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    LastName = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Email = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Phone = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    AffiliatedOrg = table.Column<string>(type: "nvarchar(max)", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Conttact", x => x.ContactId);
                });

            migrationBuilder.CreateTable(
                name: "EventSummary",
                columns: table => new
                {
                    EventSummaryId = table.Column<int>(type: "int", nullable: false)
                        .Annotation("SqlServer:Identity", "1, 1"),
                    EventId = table.Column<int>(type: "int", nullable: false),
                    Attendees = table.Column<int>(type: "int", nullable: false),
                    Challenges = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Outcomes = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Notes = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    NextSteps = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    TeamMemberID = table.Column<string>(type: "nvarchar(450)", nullable: false),
                    Status = table.Column<int>(type: "int", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_EventSummary", x => x.EventSummaryId);
                    table.ForeignKey(
                        name: "FK_EventSummary_Event_EventId",
                        column: x => x.EventId,
                        principalTable: "Event",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                    table.ForeignKey(
                        name: "FK_EventSummary_Users_TeamMemberID",
                        column: x => x.TeamMemberID,
                        principalTable: "Users",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "SchoolReport",
                columns: table => new
                {
                    SchoolReportId = table.Column<int>(type: "int", nullable: false)
                        .Annotation("SqlServer:Identity", "1, 1"),
                    Curriculum = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    LessonPlan = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    SchoolName = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    TopicsCovered = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    StudentList = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Challenges = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Successes = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Notes = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Status = table.Column<int>(type: "int", nullable: false),
                    TeamMemberId = table.Column<string>(type: "nvarchar(450)", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_SchoolReport", x => x.SchoolReportId);
                    table.ForeignKey(
                        name: "FK_SchoolReport_Users_TeamMemberId",
                        column: x => x.TeamMemberId,
                        principalTable: "Users",
                        principalColumn: "Id");
                });

            migrationBuilder.CreateIndex(
                name: "IX_EventSummary_EventId",
                table: "EventSummary",
                column: "EventId");

            migrationBuilder.CreateIndex(
                name: "IX_EventSummary_TeamMemberID",
                table: "EventSummary",
                column: "TeamMemberID");

            migrationBuilder.CreateIndex(
                name: "IX_SchoolReport_TeamMemberId",
                table: "SchoolReport",
                column: "TeamMemberId");

            migrationBuilder.AddForeignKey(
                name: "FK_ProductivityLog_GrantProgram_GrantProgramID",
                table: "ProductivityLog",
                column: "GrantProgramID",
                principalTable: "GrantProgram",
                principalColumn: "Id",
                onDelete: ReferentialAction.Cascade);
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_ProductivityLog_GrantProgram_GrantProgramID",
                table: "ProductivityLog");

            migrationBuilder.DropTable(
                name: "Conttact");

            migrationBuilder.DropTable(
                name: "EventSummary");

            migrationBuilder.DropTable(
                name: "SchoolReport");

            migrationBuilder.RenameColumn(
                name: "GrantProgramID",
                table: "ProductivityLog",
                newName: "GrantProgramId");

            migrationBuilder.RenameIndex(
                name: "IX_ProductivityLog_GrantProgramID",
                table: "ProductivityLog",
                newName: "IX_ProductivityLog_GrantProgramId");

            migrationBuilder.AlterColumn<int>(
                name: "GrantProgramId",
                table: "ProductivityLog",
                type: "int",
                nullable: true,
                oldClrType: typeof(int),
                oldType: "int");

            migrationBuilder.AddForeignKey(
                name: "FK_ProductivityLog_GrantProgram_GrantProgramId",
                table: "ProductivityLog",
                column: "GrantProgramId",
                principalTable: "GrantProgram",
                principalColumn: "Id");
        }
    }
}
