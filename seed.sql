-- Insert into EVENT table
INSERT INTO event (
    id,
    name,
    description,
    start_date,
    end_date,
    location,
    is_paid,
    total_cost,
    created_at,
    last_update
  )
VALUES (
    '97a930b8-6cde-4348-9069-c008d2d63efe',
    'Tech Conference 2024',
    'A conference about the latest in tech.',
    '2024-05-20 09:00:00+00',
    '2024-05-22 17:00:00+00',
    'Convention Center, Tech City',
    TRUE,
    299.99,
    NOW(),
    NOW()
  ),
  (
    '75726427-3b4c-417f-ad86-a569d40e5246',
    'Local Music Festival',
    'Annual music festival featuring local bands.',
    '2024-07-04 12:00:00+00',
    '2024-07-04 23:59:00+00',
    'Riverfront Park, Musicville',
    FALSE,
    0.00,
    NOW(),
    NOW()
  ) ON CONFLICT DO NOTHING;
-- Insert into ATTENDEE table
INSERT INTO attendee (id, event_id, name, created_at, last_update)
VALUES (
    'f92be039-ce59-4a13-93f6-ce3889e67100',
    '97a930b8-6cde-4348-9069-c008d2d63efe',
    'Alex Smith',
    NOW(),
    NOW()
  ),
  (
    '4c4f1fe2-276c-484b-a0ca-835b419eb9a2',
    '97a930b8-6cde-4348-9069-c008d2d63efe',
    'Jamie Doe',
    NOW(),
    NOW()
  ),
  (
    '9c938f5e-e2ff-43b3-974b-fa1dfeb0e307',
    '75726427-3b4c-417f-ad86-a569d40e5246',
    'Casey Jordan',
    NOW(),
    NOW()
  ) ON CONFLICT DO NOTHING;