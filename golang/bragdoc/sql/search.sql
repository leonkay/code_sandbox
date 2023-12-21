-- name: trackertype:list
SELECT tracker_type_id, display_tier, context_eligible, activity_key
  from b_tracker_type
  order by display_tier;

-- name: trackertype:contexteligible
SELECT tracker_type_id, display_tier, context_eligible, activity_key
  from b_tracker_type
  where context_eligible = ?
  order by activity_key;

-- name: activityevent:tier:desc
SELECT act.activity_id, act.time, act.title,
    tra.activity_key, tra.display_tier,
    ext.description
  FROM B_ACTIVITY_EVENT act
  INNER JOIN B_TRACKER_TYPE tra on act.tracker_type_id = tra.tracker_type_id
  LEFT JOIN B_ACTIVITY_EVENT_EXT ext on act.activity_id = ext.activity_id
  WHERE  tra.display_tier BETWEEN ? and ?
  ORDER BY act.time DESC;
  ;

-- name: activityevent:tier:asc
SELECT act.activity_id, act.time, act.title,
    tra.activity_key, tra.display_tier,
    ext.description
  FROM B_ACTIVITY_EVENT act
  INNER JOIN B_TRACKER_TYPE tra on act.tracker_type_id = tra.tracker_type_id
  LEFT JOIN B_ACTIVITY_EVENT_EXT ext on act.activity_id = ext.activity_id
  WHERE  tra.display_tier BETWEEN ? and ?
  ORDER BY act.time ASC;
  ;

-- name: company:byid
SELECT c.company_id, c.company_name
  FROM B_COMPANY c
  WHERE c.company_id = ?;

-- name: company:byname
SELECT c.company_id, c.company_name
  FROM B_COMPANY c
  WHERE c.company_name = ?;


-- name:title:byid
SELECT t.title_id, t.company_id, t.title, t.level
  from B_TITLE t
  WHERE t.title_id = ?;

-- name:title:byname
SELECT t.title_id, t.company_id, t.title, t.level
  from B_TITLE t
  WHERE t.company_id = ? AND t.title = ?;

-- name:title:byname:selectid
SELECT t.title_id
  from B_TITLE t
  WHERE t.company_id = ? AND t.title = ?;

-- name: insert:trackertype
INSERT INTO
  b_tracker_type (display_tier, context_eligible, activity_key)
  values (?, ?, ?);

-- name: insert:activityevent
INSERT INTO
  B_ACTIVITY_EVENT(tracker_type_id, action_id, title)
  VALUES (?, ?, ?);

-- name: insert:activityevent_ext
INSERT INTO
  B_ACTIVITY_EVENT_EXT(activity_id, 'description')
  VALUES (?, ?);
