
CREATE OR REPLACE FUNCTION register_client(
    user_name VARCHAR(50),
    pass_hash VARCHAR(255),
    email_addr VARCHAR(100),
    real_name VARCHAR(100),
    passport_number_serial VARCHAR(11),
    phone_number VARCHAR(15),
    driving_experience_ INT
) RETURNS VOID
    LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO Users (username, password_hash, email)
    VALUES (user_name, pass_hash, email_addr);

    INSERT INTO Clients (user_id, name, passport, phone, driving_experience)
    VALUES (
               (SELECT id FROM Users WHERE Users.username = user_name),
               real_name, passport_number_serial, phone_number, driving_experience_
           );
END;
$$;
