import sys
from instagramy import InstagramUser


if (len(sys.argv) != 2):
    print('Usage: python3 wilbur.py user_name')

user_name = sys.argv[1]

session_id = '51184583203%3A8YX6vOdKay4MRV%3A19'

user = InstagramUser(user_name, sessionid=session_id)

print('is user verified:', user.is_verified)

print('biography:', user.biography)

print('num posts:', user.number_of_posts)

print('num followers:', user.number_of_followers)

print('num following:', user.number_of_followings)
# print(user.user_data)