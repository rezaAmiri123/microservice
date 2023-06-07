Feature: authorize payment

  As a user I can authorize new payment

  Scenario: authorize a payment
    Given I am a registered user
    When I authorize the payment
    Then I expect the payment was authorized
