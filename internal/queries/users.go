package queries

var GetUsers = `
    SELECT
        u.id,
        u.fullname,
        u.phone,
        u.address,
        u.is_verified,
        u.subscriptions_id AS subscription_id,
        u.subscription_end_date::VARCHAR,
        u.amount_of_use
    FROM users u
    ORDER BY u.id DESC
    OFFSET $1
    LIMIT $2
`
var GetUser = `
    SELECT
        u.id,
        u.fullname,
        u.phone,
        u.address,
        u.is_verified,
        u.subscriptions_id AS subscription_id,
        u.subscription_end_date::VARCHAR,
        u.amount_of_use
    FROM users u
    WHERE u.id = $1
`
var CreateUser = `
    INSERT INTO users (
        fullname, phone, address, password, created_at, updated_at
    ) VALUES ($1, $2, $3, $4, $5, $6)
`
var CheckSubscription = `
    SELECT
        u.subscriptions_id,
        u.amount_of_use,
        (
            SELECT s.count FROM subscriptions s WHERE s.id = u.subscriptions_id
        )
    FROM users u WHERE u.id = $1
`
var BuySubscription = `
    UPDATE users SET subscriptions_id = $1, subscription_end_date = (
        SELECT NOW() + s.days FROM subscriptions s WHERE s.id = $1
    )
    WHERE id = $2
`
var UpdateUserAmountOfUse = `
    UPDATE users SET amount_of_use = amount_of_use + 1
    WHERE id = $1
`

// var CheckUserExist = "SELECT phone FROM users WHERE phone = $1"
var CheckUserExist = `
    WITH updated AS (
        UPDATE users
        SET updated_at = NOW()
        WHERE phone = $1
        AND updated_at + INTERVAL '30 seconds' <= NOW()
        RETURNING updated_at
    )
    SELECT 
        CASE 
            WHEN EXISTS (SELECT 1 FROM updated) 
            THEN 0  -- Indicates the update was performed
            ELSE EXTRACT(EPOCH FROM ((SELECT updated_at FROM users WHERE phone = $1) + INTERVAL '30 seconds') - now())::int
        END AS result;
`

var UpdateUser = `
    UPDATE users SET fullname = $1, phone = $2, address = $3,
    password = $4, updated_at = $5 WHERE id = $6
`
var UpdateUserWithoutPassword = `
    UPDATE users SET fullname = $1, phone = $2, address = $3, updated_at = $4
    WHERE id = $5
`
var VerifyUser = "UPDATE users SET is_verified = true WHERE phone = $1"
var DeleteUser = "DELETE FROM users WHERE id = $1"
var UpdateUserPassword = `
    UPDATE users u SET password = $1 WHERE u.phone = $2
    RETURNING
        u.id,
        u.fullname,
        u.phone,
        u.address,
        u.password,
        (
            SELECT
                JSON_BUILD_OBJECT(
                    'id', s.id,
                    'title', s.title,
                    'description', s.description,
                    'days', s.days::VARCHAR,
                    'count', s.count,
                    'price', s.price
                )
            FROM subscriptions s
            WHERE s.id = u.subscriptions_id
        ) AS subscription,
        u.subscription_end_date::VARCHAR,
        u.amount_of_use
`
var GetUserNotificationToken = `
    SELECT notification_token FROM users WHERE id = $1
`
var SetUserNotificationToken = `
    UPDATE users SET notification_token = $1 WHERE id = $2
`
var DeleteUserSubscription = `
    UPDATE users SET subscriptions_id = NULL, subscription_end_date = NULL,
    amount_of_use = 0 WHERE subscription_end_date <= NOW()
`
var DeleteIsNotVerifiedUsers = `
    DELETE FROM users
    WHERE is_verified = false
    AND updated_at + INTERVAL '15 minutes' <= NOW()
`

var DeleteExpiredOTPs = `
    DELETE FROM otps
    WHERE created_at + INTERVAL '30 seconds' <= NOW()
`
