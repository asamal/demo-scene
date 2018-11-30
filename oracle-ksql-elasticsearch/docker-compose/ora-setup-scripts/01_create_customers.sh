#!/bin/sh

echo 'Creating Debezium.customers table'

sqlplus Debezium/dbz@//localhost:1521/ORCLPDB1  <<- EOF

  drop table customers;
  
        create table CUSTOMERS (
                id NUMBER(10) GENERATED BY DEFAULT ON NULL AS IDENTITY (START WITH 42) NOT NULL PRIMARY KEY,
                first_name VARCHAR(50),
                last_name VARCHAR(50),
                email VARCHAR(50),
                gender VARCHAR(50),
                club_status VARCHAR(8),
                comments VARCHAR(90),
                create_ts timestamp DEFAULT CURRENT_TIMESTAMP ,
                update_ts timestamp DEFAULT CURRENT_TIMESTAMP 
        --        update_ts timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
        );
EOF

sqlplus sys/Admin123@//localhost:1521/ORCLPDB1 as sysdba <<- EOF

  ALTER TABLE debezium.customers ADD SUPPLEMENTAL LOG DATA (ALL) COLUMNS;
  GRANT SELECT ON debezium.customers to c##xstrm;

  -- From https://xanpires.wordpress.com/2013/06/26/how-to-check-the-supplemental-log-information-in-oracle/
  COLUMN LOG_GROUP_NAME HEADING 'Log Group' FORMAT A20
  COLUMN TABLE_NAME HEADING 'Table' FORMAT A20
  COLUMN ALWAYS HEADING 'Type of Log Group' FORMAT A30

  SELECT LOG_GROUP_NAME, TABLE_NAME, DECODE(ALWAYS, 'ALWAYS', 'Unconditional', NULL, 'Conditional') ALWAYS FROM DBA_LOG_GROUPS;

  exit;
EOF
