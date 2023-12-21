-- name: seed-data
INSERT OR IGNORE INTO b_tracker_type (display_tier, context_eligible, activity_key) values (100, TRUE, 'company');
INSERT OR IGNORE INTO b_tracker_type (display_tier, context_eligible, activity_key) values (100, TRUE, 'title');
INSERT OR IGNORE INTO b_tracker_type (display_tier, context_eligible, activity_key) values (100, FALSE, 'milestone');
INSERT OR IGNORE INTO b_tracker_type (display_tier, context_eligible, activity_key) values (100, FALSE, 'salary');
INSERT OR IGNORE INTO b_tracker_type (display_tier, context_eligible, activity_key) values (200, TRUE, 'role');
INSERT OR IGNORE INTO b_tracker_type (display_tier, context_eligible, activity_key) values (200, TRUE, 'project');
INSERT OR IGNORE INTO b_tracker_type (display_tier, context_eligible, activity_key) values (200, FALSE, 'yearly_review');
INSERT OR IGNORE INTO b_tracker_type (display_tier, context_eligible, activity_key) values (200, FALSE, 'achievement');
INSERT OR IGNORE INTO b_tracker_type (display_tier, context_eligible, activity_key) values (200, FALSE, 'okr');
INSERT OR IGNORE INTO b_tracker_type (display_tier, context_eligible, activity_key) values (300, FALSE, 'project_contribution');
INSERT OR IGNORE INTO b_tracker_type (display_tier, context_eligible, activity_key) values (300, FALSE, 'kpi');
INSERT OR IGNORE INTO b_tracker_type (display_tier, context_eligible, activity_key) values (300, FALSE, 'publication');
INSERT OR IGNORE INTO b_tracker_type (display_tier, context_eligible, activity_key) values (300, FALSE, 'responsibility');
INSERT OR IGNORE INTO b_tracker_type (display_tier, context_eligible, activity_key) values (400, FALSE, 'skill');
INSERT OR IGNORE INTO b_tracker_type (display_tier, context_eligible, activity_key) values (400, FALSE, 'feedback');
INSERT OR IGNORE INTO b_tracker_type (display_tier, context_eligible, activity_key) values (400, FALSE, 'presentation');
INSERT OR IGNORE INTO b_tracker_type (display_tier, context_eligible, activity_key) values (400, FALSE, 'growth');
INSERT OR IGNORE INTO b_tracker_type (display_tier, context_eligible, activity_key) values (400, FALSE, 'challenge');
INSERT OR IGNORE INTO b_tracker_type (display_tier, context_eligible, activity_key) values (500, FALSE, 'worklog');
INSERT OR IGNORE INTO b_tracker_type (display_tier, context_eligible, activity_key) values (500, FALSE, 'codecommit');
INSERT OR IGNORE INTO b_tracker_type (display_tier, context_eligible, activity_key) values (600, FALSE, 'volunteer_activity');

-- INSERT OR IGNORE INTO B_COMPANY (company_name) values ('Default Company');
-- INSERT OR IGNORE INTO B_TITLE (company_id, title, level) values (1, 'Default Title', 'L1');
