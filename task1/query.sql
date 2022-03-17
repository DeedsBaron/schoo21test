SELECT IFNULL( (SELECT DISTINCT salary
                FROM Staff
                ORDER BY salary DESC
                LIMIT 1,1), null) as SecSalary
FROM Staff
