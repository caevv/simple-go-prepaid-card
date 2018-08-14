Feature: topup card
  In order have more money to spend
  As a card owner
  I need to topup my card

  Scenario: topup card
    When I top-up for an amount of "1000"
    Then I should have a balance of "1000"
