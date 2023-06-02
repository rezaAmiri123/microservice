Feature: Register User

  Scenario: Registering a new user
    Given no user named "John Smith" exists
    When I register a new user as "John Smith"
    Then I expect the user is created
